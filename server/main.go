package main

import (
	"./api"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var config *Config

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("https://%s%s%s", config.Website.Url, config.Website.HttpsPort, r.RequestURI)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func main() {
	config = NewConfig()
	err := config.ReadFile("config.gcfg")
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	apiHandler, err := api.MakeHandler()
	if err != nil {
		log.Fatal(err)
	}
	r.PathPrefix("/api/").Handler(apiHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("../app/")))
	http.Handle("/", r)
	go func() {
		log.Printf("Serving HTTP on %s\n", config.Website.HttpPort)
		err := http.ListenAndServe(config.Website.HttpPort, http.HandlerFunc(redirectHandler))
		if err != nil {
			log.Fatal(err)
		}
	}()
	log.Printf("Serving HTTPS on %s\n", config.Website.HttpsPort)
	err = http.ListenAndServeTLS(config.Website.HttpsPort, "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatal(err)
	}
}
