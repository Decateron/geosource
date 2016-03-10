package api

import (
	"encoding/base64"
	"log"
	"net/http"
	"time"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/joshheinrichs/geosource/server/transactions"
	"github.com/joshheinrichs/geosource/server/types"
	"github.com/pborman/uuid"
)

func GetComments(w rest.ResponseWriter, req *rest.Request) {
	userId, err := GetUserId(req)
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	comments, err := transactions.GetComments(userId, req.PathParam("pid"))
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(comments)
	w.WriteHeader(http.StatusOK)
}

// func SetComment(w rest.ResponseWriter, req *rest.Request) {}

func AddComment(w rest.ResponseWriter, req *rest.Request) {
	userId, err := GetUserId(req)
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var comment types.Comment
	err = req.DecodeJsonPayload(&comment)
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	comment.Time = time.Now()
	comment.CreatorId = userId
	comment.Id = base64.RawURLEncoding.EncodeToString(uuid.NewRandom())

	err = transactions.AddComment(userId, &comment)
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func RemoveComment(w rest.ResponseWriter, req *rest.Request) {}
