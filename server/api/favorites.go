package api

import (
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/joshheinrichs/geosource/server/transactions"
)

func GetFavorites(w rest.ResponseWriter, req *rest.Request) {
	requester := GetRequesterID(req)
	favorites, err := transactions.GetFavorites(requester, req.PathParam("userID"))
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(favorites)
	w.WriteHeader(http.StatusOK)
}

func AddFavorite(w rest.ResponseWriter, req *rest.Request) {
	requester := GetRequesterID(req)
	body := struct {
		PostID string `json:"postID"`
	}{}
	err := req.DecodeJsonPayload(&body)
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println(requester, req.PathParam("userID"), body.PostID)
	err = transactions.AddFavorite(requester, req.PathParam("userID"), body.PostID)
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func RemoveFavorite(w rest.ResponseWriter, req *rest.Request) {
	requester := GetRequesterID(req)
	err := transactions.RemoveFavorite(requester, req.PathParam("userID"), req.PathParam("postID"))
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
