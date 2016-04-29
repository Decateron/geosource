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
	channels, httpErr := transactions.GetChannels(requesterID)
	if httpErr != nil {
		log.Println(httpErr)
		rest.Error(w, httpErr.Error(), httpErr.Code())
		return
	}
	w.WriteJson(channels)
	w.WriteHeader(http.StatusOK)
}

func GetChannel(w rest.ResponseWriter, req *rest.Request) {
	requesterID := GetRequesterID(req)
	channelname := req.PathParam("channelname")
	channel, httpErr := transactions.GetChannel(requesterID, channelname)
	if httpErr != nil {
		log.Println(httpErr)
		rest.Error(w, httpErr.Error(), httpErr.Code())
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
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	channel, err := types.UnmarshalChannel(jsonBody)
	if err != nil {
		log.Println("invalid channel structure", err)
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	channel.CreatorID = requesterID
	err = channel.Validate()
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	httpErr := transactions.AddChannel(channel)
	if httpErr != nil {
		log.Println(httpErr)
		rest.Error(w, httpErr.Error(), httpErr.Code())
		return
	}
	insertedChannel, httpErr := transactions.GetChannel(requesterID, channel.Name)
	if httpErr != nil {
		log.Println(httpErr)
		rest.Error(w, httpErr.Error(), httpErr.Code())
		return
	}
	w.WriteJson(insertedChannel)
	w.WriteHeader(http.StatusOK)
}

func RemoveChannel(w rest.ResponseWriter, req *rest.Request) {}
