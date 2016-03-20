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
	requester := GetRequesterID(req)
	comments, err := transactions.GetComments(requester, req.PathParam("pid"))
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
	requester := GetRequesterID(req)
	var comment types.Comment
	err := req.DecodeJsonPayload(&comment)
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	comment.Time = time.Now().UTC()
	comment.PostID = req.PathParam("pid")
	comment.CreatorID = requester
	comment.ID = base64.RawURLEncoding.EncodeToString(uuid.NewRandom())
	err = comment.Validate()
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = transactions.AddComment(requester, &comment)
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func RemoveComment(w rest.ResponseWriter, req *rest.Request) {}
