package api

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/joshheinrichs/geosource/server/transactions"
)

func GetViewers(w rest.ResponseWriter, req *rest.Request) {
	requesterID := GetRequesterID(req)
	channelname := req.PathParam("channelname")
	moderators, err := transactions.GetViewers(requesterID, channelname)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.WriteJson(moderators)
}

func AddViewer(w rest.ResponseWriter, req *rest.Request) {
	requesterID := GetRequesterID(req)
	channelname := req.PathParam("channelname")
	var body struct {
		UserID string `json:"userID"`
	}
	err := req.DecodeJsonPayload(&body)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = transactions.AddViewer(requesterID, body.UserID, channelname)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func RemoveViewer(w rest.ResponseWriter, req *rest.Request) {
	requesterID := GetRequesterID(req)
	userID := req.PathParam("userID")
	channelname := req.PathParam("channelname")
	err := transactions.RemoveViewer(requesterID, userID, channelname)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
