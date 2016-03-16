package api

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/joshheinrichs/geosource/server/transactions"
	"github.com/pborman/uuid"
)

func GetPosts(w rest.ResponseWriter, req *rest.Request) {
	userID, err := GetUserID(req)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	posts, err := transactions.GetPosts(userID)
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(posts)
	w.WriteHeader(http.StatusOK)
}

func GetPost(w rest.ResponseWriter, req *rest.Request) {
	userID, err := GetUserID(req)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	post, err := transactions.GetPost(userID, req.PathParam("pid"))
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(post)
	w.WriteHeader(http.StatusOK)
}

func AddPost(w rest.ResponseWriter, req *rest.Request) {
	userID, err := GetUserID(req)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var jsonBody json.RawMessage
	err = req.DecodeJsonPayload(&jsonBody)
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

	channel, err := transactions.GetChannel(userID, submissionChannel.Channel)
	if err != nil {
		log.Println("could not find channel ", submissionChannel.Channel)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	post, err := channel.UnmarshalSubmissionToPost(jsonBody)
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	post.CreatorID = userID
	post.ID = base64.RawURLEncoding.EncodeToString(uuid.NewRandom())
	post.Time = time.Now().UTC()

	err = post.Validate()
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = post.GenerateThumbnail()
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = transactions.AddPost(userID, post)
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func RemovePost(w rest.ResponseWriter, req *rest.Request) {}
