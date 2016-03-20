package api

import (
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/joshheinrichs/geosource/server/transactions"
)

func GetSubscriptions(w rest.ResponseWriter, req *rest.Request) {
	requester := GetRequesterID(req)
	subscriptions, err := transactions.GetSubscriptions(requester, req.PathParam("userID"))
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(subscriptions)
	w.WriteHeader(http.StatusOK)
}

func AddSubscription(w rest.ResponseWriter, req *rest.Request) {
	requester := GetRequesterID(req)
	body := struct {
		Channelname string `json:"channelname"`
	}{}
	err := req.DecodeJsonPayload(&body)
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println(requester, req.PathParam("userID"), body.Channelname)
	err = transactions.AddSubscription(requester, req.PathParam("userID"), body.Channelname)
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func RemoveSubscription(w rest.ResponseWriter, req *rest.Request) {
	requester := GetRequesterID(req)
	err := transactions.RemoveSubscription(requester, req.PathParam("userID"), req.PathParam("channelname"))
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
