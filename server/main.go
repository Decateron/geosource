package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joshheinrichs/geosource/server/api"
	"github.com/joshheinrichs/geosource/server/config"
	"github.com/joshheinrichs/geosource/server/transactions"
)

var mainConfig *config.Config

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("https://%s%s%s", mainConfig.Website.URL, mainConfig.Website.HTTPSPort, r.RequestURI)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	mainConfig = config.New()
	err := mainConfig.ReadFile("config.gcfg")
	if err != nil {
		log.Fatal(err)
	}
	api.Init(mainConfig)
	err = transactions.Init(mainConfig)
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	apiHandler, err := api.MakeHandler()
	if err != nil {
		log.Fatal(err)
	}
	r.HandleFunc("/", api.IndexHandler)
	r.PathPrefix("/api").Handler(http.StripPrefix("/api", apiHandler))
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("../app/")))
	http.Handle("/", r)
	go func() {
		log.Printf("Serving HTTP on %s\n", mainConfig.Website.HTTPPort)
		err := http.ListenAndServe(mainConfig.Website.HTTPPort, http.HandlerFunc(redirectHandler))
		if err != nil {
			log.Fatal(err)
		}
	}()
	log.Printf("Serving HTTPS on %s\n", mainConfig.Website.HTTPSPort)
	err = http.ListenAndServeTLS(mainConfig.Website.HTTPSPort, mainConfig.Website.Cert, mainConfig.Website.Key, nil)
	if err != nil {
		log.Fatal(err)
	}
}
