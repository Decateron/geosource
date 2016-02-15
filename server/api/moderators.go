package api

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/joshheinrichs/geosource/server/transactions"
)

func GetModerators(w rest.ResponseWriter, req *rest.Request) {
	// TODO: get requester
	requester := ""
	channelname := req.PathParam("channelname")
	moderators, err := transactions.GetModerators(requester, channelname)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.WriteJson(moderators)
}

func AddModerator(w rest.ResponseWriter, req *rest.Request) {
	// TODO: get requester
	requester := ""
	channelname := req.PathParam("channelname")
	var body struct {
		Username string `json:"username"`
	}
	err := req.DecodeJsonPayload(&body)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = transactions.AddModerator(requester, body.Username, channelname)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func RemoveModerator(w rest.ResponseWriter, req *rest.Request) {
	// TODO: get requester
	requester := ""
	username := req.PathParam("username")
	channelname := req.PathParam("channelname")
	err := transactions.RemoveModerator(requester, username, channelname)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
