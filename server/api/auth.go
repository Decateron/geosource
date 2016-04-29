package api

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/joshheinrichs/geosource/server/transactions"
	"github.com/joshheinrichs/geosource/server/types"
	"github.com/markbates/goth/gothic"
	"github.com/pborman/uuid"
)

// BeginAuth begins the authentication process, redirecting the user to some
// OAuth2 API depending upon the provider specified in the request path
func BeginAuth(w rest.ResponseWriter, req *rest.Request) {
	setProvider(req)
	gothic.BeginAuthHandler(w.(http.ResponseWriter), req.Request)
}

// CallbackAuth handles callbacks from OAuth 2 APIs, signing in users and
// creating them if they do not exist. Once the user is signed in, their userID
// is stored into their session for identification. Afterwards they are
// redirected to the main site.
func CallbackAuth(w rest.ResponseWriter, req *rest.Request) {
	setProvider(req)
	gothUser, err := gothic.CompleteUserAuth(w.(http.ResponseWriter), req.Request)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user, httpErr := transactions.GetUserByEmail(gothUser.Email)
	if httpErr != nil {
		rest.Error(w, httpErr.Error(), httpErr.Code())
		return
	}
	if user == nil {
		user = &types.User{
			Name:   gothUser.Name,
			Avatar: gothUser.AvatarURL,
			Email:  gothUser.Email,
			ID:     base64.RawURLEncoding.EncodeToString(uuid.NewRandom()),
		}
		httpErr = transactions.AddUser(user)
		if httpErr != nil {
			rest.Error(w, httpErr.Error(), httpErr.Code())
			return
		}
	}
	session, err := gothic.Store.Get(req.Request, gothic.SessionName)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["userID"] = user.ID
	err = session.Save(req.Request, w.(http.ResponseWriter))
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w.(http.ResponseWriter), req.Request, fmt.Sprintf("https://%s%s", apiConfig.Website.URL, apiConfig.Website.HTTPSPort), http.StatusTemporaryRedirect)
}

// Logout Logs out the user, deleting their associated user data from their
// session.
func Logout(w rest.ResponseWriter, req *rest.Request) {
	session, err := gothic.Store.Get(req.Request, gothic.SessionName)
	if err != nil {
		log.Printf("logout attempted by user that was not logged in.\n")
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	delete(session.Values, "userID")
	err = session.Save(req.Request, w.(http.ResponseWriter))
	if err != nil {
		log.Printf("error saving session.\n")
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w.(http.ResponseWriter), req.Request, fmt.Sprintf("https://%s%s", apiConfig.Website.URL, apiConfig.Website.HTTPSPort), http.StatusTemporaryRedirect)
}

// setProvider adds the provider path parameter from the given rest quest as a
// query value to the request URL. Gothic requires the provider as a query
// value, so this fixes that incompatibility error.
func setProvider(req *rest.Request) {
	v := req.Request.URL.Query()
	v.Set("provider", req.PathParam("provider"))
	req.Request.URL.RawQuery = v.Encode()
}
