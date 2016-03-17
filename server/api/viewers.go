package api

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/joshheinrichs/geosource/server/transactions"
)

func GetViewers(w rest.ResponseWriter, req *rest.Request) {
	requester, err := GetUserID(req)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	channelname := req.PathParam("channelname")
	moderators, err := transactions.GetViewers(requester, channelname)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.WriteJson(moderators)
}

func AddViewer(w rest.ResponseWriter, req *rest.Request) {
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
	err = transactions.AddViewer(requester, body.UserID, channelname)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func RemoveViewer(w rest.ResponseWriter, req *rest.Request) {
	requester, err := GetUserID(req)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userID := req.PathParam("userID")
	channelname := req.PathParam("channelname")
	err = transactions.RemoveViewer(requester, userID, channelname)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
