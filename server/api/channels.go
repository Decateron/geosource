package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/joshheinrichs/geosource/server/types"
	"github.com/markbates/goth/gothic"
)

func GetChannels(w rest.ResponseWriter, req *rest.Request) {}

func GetChannel(w rest.ResponseWriter, req *rest.Request) {}

// func SetChannel(w rest.ResponseWriter, req *rest.Request) {}

func AddChannel(w rest.ResponseWriter, req *rest.Request) {
	session, err := gothic.Store.Get(req.Request, gothic.SessionName)
	if err != nil {
		log.Println("channel creation attempted by user that was not logged in")
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userId, ok := session.Values["userid"].(string)
	if !ok {
		log.Println("invalid user id cookie")
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var jsonBody json.RawMessage
	err = req.DecodeJsonPayload(&jsonBody)
	if err != nil {
		log.Println("invalid json body")
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	channel, err := types.UnmarshalChannel(jsonBody)
	if err != nil {
		log.Println("invalid channel structure", err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	channel.CreatorId = userId
	data, err := json.Marshal(channel)
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println(string(data))
	w.WriteHeader(http.StatusOK)
}

func RemoveChannel(w rest.ResponseWriter, req *rest.Request) {}
