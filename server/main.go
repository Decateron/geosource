package main

import (
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    r.PathPrefix("/").Handler(http.FileServer(http.Dir("../app/")))
    http.Handle("/", r)
    http.ListenAndServe(":8000", nil)
}
