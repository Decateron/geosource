package api

import (
	"../transaction"
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
)

func GetViewers(w rest.ResponseWriter, req *rest.Request) {
	// TODO: get requester
	requester := ""
	channelname := req.PathParam("channelname")
	moderators, err := transaction.GetViewers(requester, channelname)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.WriteJson(moderators)
}

func AddViewer(w rest.ResponseWriter, req *rest.Request) {
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
	err = transaction.AddViewer(requester, body.Username, channelname)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func RemoveViewer(w rest.ResponseWriter, req *rest.Request) {
	// TODO: get requester
	requester := ""
	username := req.PathParam("username")
	channelname := req.PathParam("channelname")
	err := transaction.RemoveViewer(requester, username, channelname)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
