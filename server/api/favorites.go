package api

import (
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/joshheinrichs/geosource/server/transactions"
)

func GetFavorites(w rest.ResponseWriter, req *rest.Request) {
	requesterID := GetRequesterID(req)
	favorites, err := transactions.GetFavorites(requesterID, req.PathParam("userID"))
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(favorites)
	w.WriteHeader(http.StatusOK)
}

func AddFavorite(w rest.ResponseWriter, req *rest.Request) {
	requesterID := GetRequesterID(req)
	body := struct {
		PostID string `json:"postID"`
	}{}
	err := req.DecodeJsonPayload(&body)
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println(requesterID, req.PathParam("userID"), body.PostID)
	err = transactions.AddFavorite(requesterID, req.PathParam("userID"), body.PostID)
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func RemoveFavorite(w rest.ResponseWriter, req *rest.Request) {
	requesterID := GetRequesterID(req)
	err := transactions.RemoveFavorite(requesterID, req.PathParam("userID"), req.PathParam("postID"))
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
