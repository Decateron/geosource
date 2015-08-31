package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"log"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
)
// "github.com/google/google-api-go-client/plus/v1"


// config is the configuration specification supplied to the OAuth package.
var oauthConfig oauth2.Config

// store initializes the Gorilla session store.
var store = sessions.NewCookieStore([]byte(randomString(32)))

// Token represents an OAuth token response.
type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	IdToken     string `json:"id_token"`
}

// ClaimSet represents an IdToken response.
type ClaimSet struct {
	Sub string
}

type UserInfo struct {
    Email string `json:"email"`
}

// exchange takes an authentication code and exchanges it with the OAuth
// endpoint for a Google API bearer token and a Google+ ID
func exchange(code string) (accessToken string, idToken string, err error) {
	tok, err := oauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return "", "", fmt.Errorf("Error while exchanging code: %v", err)
	}

	// TODO: return ID token in second parameter from updated oauth2 interface
	return tok.AccessToken, tok.Extra("id_token").(string), nil
}

// decodeIdToken takes an ID Token and decodes it to fetch the Google+ ID within
func decodeIdToken(idToken string) (gplusID string, err error) {
	// An ID token is a cryptographically-signed JSON object encoded in base 64.
	// Normally, it is critical that you validate an ID token before you use it,
	// but since you are communicating directly with Google over an
	// intermediary-free HTTPS channel and using your Client Secret to
	// authenticate yourself to Google, you can be confident that the token you
	// receive really comes from Google and is valid. If your server passes the ID
	// token to other components of your app, it is extremely important that the
	// other components validate the token before using it.
	var set ClaimSet
	if idToken != "" {
		// Check that the padding is correct for a base64decode
		parts := strings.Split(idToken, ".")
		if len(parts) < 2 {
			return "", fmt.Errorf("Malformed ID token")
		}
		// Decode the ID token
		b, err := base64Decode(parts[1])
		if err != nil {
			return "", fmt.Errorf("Malformed ID token: %v", err)
		}
		err = json.Unmarshal(b, &set)
		if err != nil {
			return "", fmt.Errorf("Malformed ID token: %v", err)
		}
	}
	return set.Sub, nil
}

// randomString returns a random string with the specified length
func randomString(length int) (str string) {
	b := make([]byte, length)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func base64Decode(s string) ([]byte, error) {
	// add back missing padding
	switch len(s) % 4 {
	case 2:
		s += "=="
	case 3:
		s += "="
	}
	return base64.URLEncoding.DecodeString(s)
}

func getUserInfo(accessToken string) (*UserInfo, error) {
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + accessToken)
	if err != nil {
		return nil, err
	}

    defer response.Body.Close()
    var userInfo UserInfo
    err = json.NewDecoder(response.Body).Decode(&userInfo)
    if err != nil {
    	return nil, err
    }
    return &userInfo, nil
}

// connect exchanges the one-time authorization code for a token and stores the
// token in the session
func login(w http.ResponseWriter, r *http.Request) {
	type RequestBody struct {
		Code *string `json:"code"`
	}
	type ResponseBody struct {
		Username *string `json:"username"`
	}

	// Ensure that the request is not a forgery and that the user sending this
	// connect request is the expected user
	session, err := store.Get(r, "sessionName")
	if err != nil {
		log.Println("Error:", err)
		json.NewEncoder(w).Encode(Error{"Error fetching session"})
		return
	}
	if r.FormValue("state") != session.Values["state"].(string) {
		log.Println("Error:", err)
		json.NewEncoder(w).Encode(Error{"Invalid state parameter"})
		return
	}

	defer r.Body.Close()
	var requestBody RequestBody
	err =  json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		log.Println("Error:", err)
		json.NewEncoder(w).Encode(Error{"Error parsing request body"})
		return
	}

	fmt.Println(requestBody.Code)

	if requestBody.Code == nil && session.Values["acessToken"] != nil {
		if session.Values["username"] != nil {
			username := session.Values["username"].(string)
			json.NewEncoder(w).Encode(ResponseBody{&username})
		} else {
			json.NewEncoder(w).Encode(ResponseBody{nil})
		}
		return
	} else if requestBody.Code == nil {
		json.NewEncoder(w).Encode(ResponseBody{nil})
		return
	}

	accessToken, idToken, err := exchange(*requestBody.Code)
	if err != nil {
		log.Println("Error:", err)
		json.NewEncoder(w).Encode(Error{"Error exchanging code for access token"})
		return //&appError{err, "Error exchanging code for access token", 500}
	}
	gplusID, err := decodeIdToken(idToken)
	if err != nil {
		json.NewEncoder(w).Encode(Error{"Error decoding ID token"})
		return
	}

	// Check if the user is already connected
	storedToken := session.Values["accessToken"]
	storedGPlusID := session.Values["gplusID"]
	if storedToken != nil && storedGPlusID == gplusID {
		log.Println("User already connected")
		if session.Values["username"] != nil {
			username := session.Values["username"].(string)
			json.NewEncoder(w).Encode(ResponseBody{&username})
		} else {
			json.NewEncoder(w).Encode(ResponseBody{nil})
		}
		return
	}
	session.Values["accessToken"] = accessToken
	session.Values["gplusID"] = gplusID
	
	userInfo, err := getUserInfo(accessToken)
	if err != nil {
		log.Println("Error:", err)
		json.NewEncoder(w).Encode(Error{"Internal server error"})
		return
	}
	session.Values["email"] = userInfo.Email

	username, err := getUsername(userInfo.Email)
	if err != nil {
		log.Println("Error:", err)
		json.NewEncoder(w).Encode(Error{"Internal server error"})
		return
	} else if username != nil {
		session.Values["username"] = username
	}
	
	err = session.Save(r, w)
	if err != nil {
		log.Println("Error:", err)
		json.NewEncoder(w).Encode(Error{"Internal server error"})
		return
	}

	json.NewEncoder(w).Encode(ResponseBody{username})
}

func logout(w http.ResponseWriter, r *http.Request) {
	type ResponseBody struct {}

	// Only disconnect a connected user
	session, err := store.Get(r, "sessionName")
	if err != nil {
		json.NewEncoder(w).Encode(Error{"Error fetching session"})
		return
	}

	// Reset the user's session
	session.Values["accessToken"] = nil
	err = session.Save(r, w)
	if err != nil {
		json.NewEncoder(w).Encode(Error{"Error saving session"})
		return
	}

	json.NewEncoder(w).Encode(ResponseBody{})
}

// disconnect revokes the current user's token and resets their session
func disconnect(w http.ResponseWriter, r *http.Request) {
	type ResponseBody struct {}

	// Only disconnect a connected user
	session, err := store.Get(r, "sessionName")
	if err != nil {
		json.NewEncoder(w).Encode(Error{"Error fetching session"})
		return
	}
	token := session.Values["accessToken"]
	if token == nil {
		json.NewEncoder(w).Encode(Error{"Current user not connected"})
		return
	}

	// Execute HTTP GET request to revoke current token
	url := "https://accounts.google.com/o/oauth2/revoke?token=" + token.(string)
	_, err = http.Get(url)
	if err != nil {
		json.NewEncoder(w).Encode(Error{"Failed to revoke token for a given user"})
		return
	}

	// Reset the user's session
	session.Values["accessToken"] = nil
	err = session.Save(r, w)
	if err != nil {
		json.NewEncoder(w).Encode(Error{"Error saving session"})
		return
	}

	
	json.NewEncoder(w).Encode(ResponseBody{})
}