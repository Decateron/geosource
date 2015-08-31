package main

import (
	"database/sql"
	"regexp"
	"errors"
)

func addUser(email, username string) error {
	matched, err := regexp.MatchString("^.*@.*$", email)
	if err != nil {
		return errors.New("Internal server error")
	} else if !matched {
		return errors.New("Invalid email")
	} else if len(email) > 254 {
		return errors.New("Emails must be 254 characters or less in length")
	}

	emailUsername, err := getUsername(email)
	if err != nil {
		return errors.New("Internal server error")
	} else if emailUsername != nil {
		return errors.New("You have already created a username")
	}

	matched, err = regexp.MatchString("^[a-zA-Z0-9]+$", username)
	if err != nil {
		return errors.New("Internal server error")
	} else if !matched {
		return errors.New("Usernames must consist of alphanumeric characters")
	} else if len(username) < 3 || len(username) > 20  {
		return errors.New("Usernames must be between 3 and 20 characters in length")
	}

	usernameEmail, err := getEmail(username)
	if err != nil {
		return errors.New("Internal server error")
	} else if usernameEmail != nil {
		return errors.New("A user with that username already exists")
	}

	_, err = db.Exec("INSERT INTO users (u_email, u_username) VALUES ($1, $2)", email, username)
	if err != nil {
		return errors.New("Internal server error")
	}
	return nil
}

func getUsername(email string) (*string, error) {
	var username string
	err := db.QueryRow("SELECT u_username FROM users WHERE u_email = $1", email).Scan(&username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	} 
	return &username, nil
}

func getEmail(username string) (*string, error) {
	var email string
	err := db.QueryRow("SELECT u_email FROM users WHERE u_username = $1", username).Scan(&email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	} 
	return &email, nil
}