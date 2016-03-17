package api

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/joshheinrichs/geosource/server/transactions"
)

func GetBans(w rest.ResponseWriter, req *rest.Request) {
	requester, err := GetUserID(req)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	channelname := req.PathParam("channelname")
	moderators, err := transactions.GetBans(requester, channelname)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.WriteJson(moderators)
}

func AddBan(w rest.ResponseWriter, req *rest.Request) {
	requester, err := GetUserID(req)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	channelname := req.PathParam("channelname")
	var body struct {
		UserID string `json:"userID"`
	}
	err = req.DecodeJsonPayload(&body)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = transactions.AddBan(requester, body.UserID, channelname)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func RemoveBan(w rest.ResponseWriter, req *rest.Request) {
	requester, err := GetUserID(req)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userID := req.PathParam("userID")
	channelname := req.PathParam("channelname")
	err = transactions.RemoveBan(requester, userID, channelname)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
