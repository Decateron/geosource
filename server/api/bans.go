package api

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/joshheinrichs/geosource/server/transactions"
)

func GetBans(w rest.ResponseWriter, req *rest.Request) {
	requesterID := GetRequesterID(req)
	channelname := req.PathParam("channelname")
	moderators, err := transactions.GetBans(requesterID, channelname)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.WriteJson(moderators)
}

func AddBan(w rest.ResponseWriter, req *rest.Request) {
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
	httpErr := transactions.AddBan(requesterID, body.UserID, channelname)
	if httpErr != nil {
		rest.Error(w, httpErr.Error(), httpErr.Code())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func RemoveBan(w rest.ResponseWriter, req *rest.Request) {
	requesterID := GetRequesterID(req)
	userID := req.PathParam("userID")
	channelname := req.PathParam("channelname")
	err := transactions.RemoveBan(requesterID, userID, channelname)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
