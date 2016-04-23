package api

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/gorilla/schema"
	"github.com/joshheinrichs/geosource/server/transactions"
	"github.com/pborman/uuid"
)

func GetPosts(w rest.ResponseWriter, req *rest.Request) {
	requesterID := GetRequesterID(req)
	log.Println(req.URL.Query().Encode())

	var postQueryParams transactions.PostQueryParams
	decoder := schema.NewDecoder()
	err := decoder.Decode(&postQueryParams, req.URL.Query())
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println(postQueryParams)

	posts, err := transactions.GetPosts(requesterID, &postQueryParams)
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteJson(posts)
	w.WriteHeader(http.StatusOK)
}

func GetPost(w rest.ResponseWriter, req *rest.Request) {
	requesterID := GetRequesterID(req)
	post, err := transactions.GetPost(requesterID, req.PathParam("pid"))
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(post)
	w.WriteHeader(http.StatusOK)
}

func AddPost(w rest.ResponseWriter, req *rest.Request) {
	requesterID := GetRequesterID(req)
	var jsonBody json.RawMessage
	err := req.DecodeJsonPayload(&jsonBody)
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	submissionChannel := struct {
		Channel string `json:"channel"`
	}{}
	err = json.Unmarshal(jsonBody, &submissionChannel)
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	channel, err := transactions.GetChannel(requesterID, submissionChannel.Channel)
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	post, err := channel.UnmarshalSubmissionToPost(jsonBody)
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = post.Validate()
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	post.CreatorID = requesterID
	post.ID = base64.RawURLEncoding.EncodeToString(uuid.NewRandom())
	post.Time = time.Now().UTC()
	err = post.GenerateThumbnail()
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = transactions.AddPost(requesterID, post)
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	insertedPost, err := transactions.GetPost(requesterID, post.ID)
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(insertedPost)
	w.WriteHeader(http.StatusOK)
}

func RemovePost(w rest.ResponseWriter, req *rest.Request) {}
