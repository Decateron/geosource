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
	requesterID := GetRequesterID(req)
	channels, err := transactions.GetChannels(requesterID)
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(channels)
	w.WriteHeader(http.StatusOK)
}

func GetChannel(w rest.ResponseWriter, req *rest.Request) {
	requesterID := GetRequesterID(req)
	channelname := req.PathParam("channelname")
	channel, err := transactions.GetChannel(requesterID, channelname)
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
	requesterID := GetRequesterID(req)
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
	channel.CreatorID = requesterID
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
	insertedChannel, err := transactions.GetChannel(requesterID, channel.Name)
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(insertedChannel)
	w.WriteHeader(http.StatusOK)
}

func RemoveChannel(w rest.ResponseWriter, req *rest.Request) {}
