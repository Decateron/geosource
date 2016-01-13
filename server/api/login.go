package api

import (
	"../transaction"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	"net/http"
	"strings"
)

// "github.com/google/google-api-go-client/plus/v1"

const (
	SESSION_NAME = "sessionName"
)

var ErrFetchingSession error = errors.New("Error fetching session.")

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
func Login(w rest.ResponseWriter, req *rest.Request) {
	type RequestBody struct {
		Code *string `json:"code"`
	}
	type ResponseBody struct {
		Username *string `json:"username"`
	}

	// Ensure that the request is not a forgery and that the user sending this
	// connect request is the expected user
	session, err := store.Get(req.Request, SESSION_NAME)
	if err != nil {
		rest.Error(w, "Error fetching session", http.StatusBadRequest)
		return
	}
	if req.FormValue("state") != session.Values["state"].(string) {
		rest.Error(w, "Invalid state parameter", http.StatusBadRequest)
		return
	}

	defer req.Body.Close()
	var requestBody RequestBody
	err = req.DecodeJsonPayload(&requestBody)
	if err != nil {
		rest.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	fmt.Println(requestBody.Code)

	if requestBody.Code == nil && session.Values["acessToken"] != nil {
		if session.Values["username"] != nil {
			username := session.Values["username"].(string)
			w.WriteJson(ResponseBody{&username})
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteJson(ResponseBody{nil})
			w.WriteHeader(http.StatusOK)
		}
		return
	} else if requestBody.Code == nil {
		w.WriteJson(ResponseBody{nil})
		w.WriteHeader(http.StatusOK)
		return
	}

	accessToken, idToken, err := exchange(*requestBody.Code)
	if err != nil {
		rest.Error(w, "Error exchanging code for access token", http.StatusBadRequest)
		return
	}
	gplusID, err := decodeIdToken(idToken)
	if err != nil {
		rest.Error(w, "Error decoding ID token", http.StatusBadRequest)
		return
	}

	// Check if the user is already connected
	storedToken := session.Values["accessToken"]
	storedGPlusID := session.Values["gplusID"]
	if storedToken != nil && storedGPlusID == gplusID {
		if session.Values["username"] != nil {
			username := session.Values["username"].(string)
			w.WriteJson(ResponseBody{&username})
		} else {
			w.WriteJson(ResponseBody{nil})
		}
		w.WriteHeader(http.StatusOK)
		return
	}
	session.Values["accessToken"] = accessToken
	session.Values["gplusID"] = gplusID

	userInfo, err := getUserInfo(accessToken)
	if err != nil {
		rest.Error(w, "Internal server error", http.StatusBadRequest)
		return
	}
	session.Values["email"] = userInfo.Email

	username, err := transaction.GetUsername(userInfo.Email)
	if err != nil {
		rest.Error(w, "Internal server error", http.StatusBadRequest)
		return
	} else if username != nil {
		session.Values["username"] = username
	}

	err = session.Save(req.Request, w.(http.ResponseWriter))
	if err != nil {
		rest.Error(w, "Internal server error", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.WriteJson(ResponseBody{username})
}

func Logout(w rest.ResponseWriter, req *rest.Request) {
	// Only disconnect a connected user
	session, err := store.Get(req.Request, SESSION_NAME)
	if err != nil {
		rest.Error(w, "Error fetching session", http.StatusBadRequest)
		return
	}

	// Reset the user's session
	session.Values["accessToken"] = nil
	err = session.Save(req.Request, w.(http.ResponseWriter))
	if err != nil {
		rest.Error(w, "Error saving session", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Disconnect revokes the current user's token and resets their session
func Disconnect(w rest.ResponseWriter, req *rest.Request) {
	// Only disconnect a connected user
	session, err := store.Get(req.Request, SESSION_NAME)
	if err != nil {
		rest.Error(w, "Error fetching session", http.StatusBadRequest)
		return
	}
	token := session.Values["accessToken"]
	if token == nil {
		rest.Error(w, "Current user not connected", http.StatusBadRequest)
		return
	}

	// Execute HTTP GET request to revoke current token
	url := "https://accounts.google.com/o/oauth2/revoke?token=" + token.(string)
	_, err = http.Get(url)
	if err != nil {
		rest.Error(w, "Failed to revoke token for a given user", http.StatusBadRequest)
		return
	}

	// Reset the user's session
	session.Values["accessToken"] = nil
	err = session.Save(req.Request, w.(http.ResponseWriter))
	if err != nil {
		rest.Error(w, "Error saving session", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
