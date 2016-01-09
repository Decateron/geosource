package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var config *Config

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://"+config.Website.Url+config.Website.HttpsPort+r.RequestURI, http.StatusTemporaryRedirect)
}

func main() {
	config := NewConfig()
	err := config.ReadFile("config.gcfg")
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("../app/")))
	http.Handle("/", r)
	go func() {
		err := http.ListenAndServe(config.Website.HttpPort, http.HandlerFunc(redirectHandler))
		if err != nil {
			log.Fatal(err)
		}
	}()
	err = http.ListenAndServeTLS(config.Website.HttpsPort, "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatal(err)
	}
}
