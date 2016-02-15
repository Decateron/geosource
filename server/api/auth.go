package api

import (
	"code.google.com/p/go-uuid/uuid"
	"encoding/base64"
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/joshheinrichs/geosource/server/transaction"
	"github.com/joshheinrichs/geosource/server/types"
	"github.com/markbates/goth/gothic"
	"log"
	"net/http"
)

func BeginAuth(w rest.ResponseWriter, req *rest.Request) {
	setProvider(req)
	gothic.BeginAuthHandler(w.(http.ResponseWriter), req.Request)
}

func CallbackAuth(w rest.ResponseWriter, req *rest.Request) {
	setProvider(req)
	gothUser, err := gothic.CompleteUserAuth(w.(http.ResponseWriter), req.Request)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user, err := transaction.GetUserByEmail(gothUser.Email)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if user == nil {
		user = &types.User{
			Name:   gothUser.Name,
			Avatar: gothUser.AvatarURL,
			Email:  gothUser.Email,
			Id:     base64.RawURLEncoding.EncodeToString(uuid.NewRandom()),
		}
		err = transaction.AddUser(user)
		if err != nil {
			rest.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	session, err := gothic.Store.Get(req.Request, gothic.SessionName)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["userid"] = user.Id
	err = session.Save(req.Request, w.(http.ResponseWriter))
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w.(http.ResponseWriter), req.Request, "https://localhost:8080/", http.StatusTemporaryRedirect)
}

func Logout(w rest.ResponseWriter, req *rest.Request) {
	session, err := gothic.Store.Get(req.Request, gothic.SessionName)
	if err != nil {
		log.Printf("logout attempted by user that was not logged in.\n")
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["userid"] = ""
	err = session.Save(req.Request, w.(http.ResponseWriter))
	if err != nil {
		log.Printf("error saving session.\n")
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w.(http.ResponseWriter), req.Request, "https://localhost:8080/", http.StatusTemporaryRedirect)
}

// Adds the provider path parameter from the given rest quest as a query value
// to the request URL. Gothic requires the provider as a query value, so this
// fixes that incompatibility error
func setProvider(req *rest.Request) {
	v := req.Request.URL.Query()
	v.Set("provider", req.PathParam("provider"))
	req.Request.URL.RawQuery = v.Encode()
}
