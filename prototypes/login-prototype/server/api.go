package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func postUser(w http.ResponseWriter, r *http.Request) {
	type RequestBody struct {
		Username string `json:"username"`
	}
	type ResponseBody struct {
		Username string `json:"username"`
	}

	session, err := store.Get(r, "sessionName")
	if err != nil {
		json.NewEncoder(w).Encode(Error{"Error fetching session"})
		return
	}

	if session.Values["accessToken"] == nil {
		json.NewEncoder(w).Encode(Error{"You must log in to perform this action"})
		return
	}

	email := session.Values["email"].(string)

	defer r.Body.Close()
	var requestBody RequestBody
	err = json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		json.NewEncoder(w).Encode(Error{"Error parsing request body"})
		fmt.Println("Error:", err)
		return
	}

	err = addUser(email, requestBody.Username)
	if err != nil {
		json.NewEncoder(w).Encode(Error{err.Error()})
		fmt.Println("Error:", err)
		return
	}

	session.Values["username"] = requestBody.Username
	err = session.Save(r, w)
	if err != nil {
		fmt.Println("Error:", err)
		json.NewEncoder(w).Encode(Error{"Error saving session"})
		return
	}

	json.NewEncoder(w).Encode(ResponseBody{requestBody.Username})
}
