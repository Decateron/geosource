package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"database/sql"
	_ "github.com/lib/pq"

    "github.com/gorilla/mux"
    "code.google.com/p/gcfg"
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
)

type Config struct {
	Google struct {
		ClientID     string
		ClientSecret string
	}
	Postgresql struct {
		Host     string
		Database string
		User     string
		Password string
	}
}

type Error struct {
	Error string `json:"error"`
}

var config Config
var db *sql.DB 

// indexTemplate is the HTML template we use to present the index page.
var indexTemplate = template.Must(template.ParseFiles("app/index.html"))

// index sets up a session for the current user and serves the index page
func index(w http.ResponseWriter, r *http.Request) {
	// Create a state token to prevent request forgery and store it in the session
	// for later validation
	session, err := store.Get(r, "sessionName")
	if err != nil {
		log.Println("error fetching session:", err)
		// Ignore the initial session fetch error, as Get() always returns a
		// session, even if empty.
		// return
	}
	state := randomString(64)
	fmt.Println(state)
	session.Values["state"] = state
	session.Save(r, w)

	// Fill in the missing fields in index.html
	var data = struct {
		ClientId, State string
	}{config.Google.ClientID, state}

	// Render and serve the HTML
	err = indexTemplate.Execute(w, data)
	if err != nil {
		log.Println("error rendering template:", err)
		return
	}
}

func loginSetup() error {
	err := gcfg.ReadFileInto(&config, "config.gcfg")
	if err != nil {
		return err
	}

	oauthConfig = oauth2.Config{
		ClientID:     config.Google.ClientID,
		ClientSecret: config.Google.ClientSecret,
		// Scope determines which API calls you are authorized to make
		Scopes:   []string{"https://www.googleapis.com/auth/plus.login", "https://www.googleapis.com/auth/userinfo.email"},
		Endpoint: google.Endpoint,
		// Use "postmessage" for the code-flow for server side apps
		RedirectURL: "postmessage",
	}
	return nil
}

func databaseSetup() (err error) {
	login := fmt.Sprintf("host=%s dbname=%s user=%s password=%s",
		config.Postgresql.Host, 
		config.Postgresql.Database, 
		config.Postgresql.User, 
		config.Postgresql.Password)

	db, err = sql.Open("postgres", login)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	
	err := loginSetup()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = databaseSetup()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	
    r := mux.NewRouter()
    r.HandleFunc("/", index)
    r.HandleFunc("/api/login", login)
    r.HandleFunc("/api/logout", logout)
    r.HandleFunc("/api/users", postUser).Methods("POST")

    r.PathPrefix("/").Handler(http.FileServer(http.Dir("app/")))

    http.Handle("/", r)
    err = http.ListenAndServe(":8000", nil)
    if err != nil {
		log.Fatal(err)
	}
}
