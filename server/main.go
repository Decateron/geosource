package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	WEBSITE_URL = "localhost"
	HTTP_PORT   = ":8000"
	HTTPS_PORT  = ":8080"
)

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://"+WEBSITE_URL+HTTPS_PORT+r.RequestURI, http.StatusMovedPermanently)
}

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("../app/")))
	http.Handle("/", r)
	go func() {
		err := http.ListenAndServe(HTTP_PORT, http.HandlerFunc(redirectHandler))
		if err != nil {
			log.Fatal(err)
		}
	}()
	err := http.ListenAndServeTLS(HTTPS_PORT, "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatal(err)
	}
}
