package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/joshheinrichs/geosource/server/transactions"
	"github.com/joshheinrichs/geosource/server/types"
)

func GetChannels(w rest.ResponseWriter, req *rest.Request) {
	requester := GetRequesterID(req)
	channels, err := transactions.GetChannels(requester)
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(channels)
	w.WriteHeader(http.StatusOK)
}

func GetChannel(w rest.ResponseWriter, req *rest.Request) {
	requester := GetRequesterID(req)
	channelname := req.PathParam("channelname")
	channel, err := transactions.GetChannel(requester, channelname)
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(channel)
	w.WriteHeader(http.StatusOK)
}

// func SetChannel(w rest.ResponseWriter, req *rest.Request) {}

func AddChannel(w rest.ResponseWriter, req *rest.Request) {
	requester := GetRequesterID(req)
	var jsonBody json.RawMessage
	err := req.DecodeJsonPayload(&jsonBody)
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
	channel.CreatorID = requester
	err = channel.Validate()
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = transactions.AddChannel(channel)
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	insertedChannel, err := transactions.GetChannel(requester, channel.Name)
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(insertedChannel)
	w.WriteHeader(http.StatusOK)
}

func RemoveChannel(w rest.ResponseWriter, req *rest.Request) {}
