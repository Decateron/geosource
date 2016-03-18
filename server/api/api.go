package api

import (
	"errors"
	"html/template"
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/joshheinrichs/geosource/server/config"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/gplus"
)

var apiConfig *config.Config

var ErrInsufficientPermission error = errors.New("Insufficient permission.")
var indexTemplate = template.Must(template.ParseFiles("../app/index.html"))

func Init(config *config.Config) {
	apiConfig = config
	rest.ErrorFieldName = "error"
	goth.UseProviders(
		gplus.New(config.Google.ClientID, config.Google.ClientSecret, config.Google.CallbackUrl),
		// facebook.New(config.Facebook.ClientID, config.Facebook.ClientSecret, config.Facebook.Callback),
		// twitter.New(config.Twitter.ClientID, config.Twitter.ClientSecret, config.Twitter.Callback),
	)
}

func MakeHandler() (http.Handler, error) {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		// Auth
		rest.Get("/auth/#provider", BeginAuth),
		rest.Get("/auth/#provider/callback", CallbackAuth),
		rest.Post("/logout", Logout),

		// Channels
		rest.Get("/channels", GetChannels),
		rest.Get("/channels/#channelname", GetChannel),
		// rest.Put("/channels/#channelname", SetChannel),
		rest.Post("/channels", AddChannel),
		rest.Delete("/channels/#channelname", RemoveChannel),

		// Moderators
		rest.Get("/channels/#channelname/moderators", GetModerators),
		rest.Post("/channels/#channelname/moderators", AddModerator),
		rest.Delete("/channels/#channelname/moderators/#userID", RemoveModerator),

		// Viewers
		rest.Get("/channels/#channelname/viewers", GetViewers),
		rest.Post("/channels/#channelname/viewers", AddViewer),
		rest.Delete("/channels/#channelname/viewers/#userID", RemoveViewer),

		// Banned
		rest.Get("/channels/#channelname/bans", GetBans),
		rest.Post("/channels/#channelname/bans", AddBan),
		rest.Delete("/channels/#channelname/bans/#userID", RemoveBan),

		// Posts
		rest.Get("/posts", GetPosts),
		rest.Get("/posts/#pid", GetPost),
		// rest.Put("/posts/#pid", SetPost),
		rest.Post("/posts", AddPost),
		rest.Delete("/posts/#pid", RemovePost),

		// Comments
		rest.Get("/posts/#pid/comments", GetComments),
		// rest.Put("/posts/#pid/comments/#cid", SetComment),
		rest.Post("/posts/#pid/comments", AddComment),
		rest.Delete("/posts/#pid/comments/#cid", RemoveComment),

		// Users
		rest.Get("/users", GetUsers),
		rest.Get("/users/#userID", GetUser),
		// rest.Put("/users/#userID", SetUser),
		rest.Post("/users/#userID", AddUser),
		rest.Delete("/users/#userID", RemoveUser),

		// Subscriptions
		rest.Get("/users/#userID/subscriptions", GetSubscriptions),
		rest.Post("/users/#userID/subscriptions", AddSubscription),
		rest.Delete("/users/#userID/subscriptions/#channelname", RemoveSubscription),

		// Favorites
		rest.Get("/users/#userID/favorites", GetFavorites),
		rest.Post("/users/#userID/favorites", AddFavorite),
		rest.Delete("/users/#userID/favorites/#postID", RemoveFavorite),
	)
	if err != nil {
		return nil, err
	}
	api.SetApp(router)
	return api.MakeHandler(), nil
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// session, _ := store.Get(r, gothic.SessionName)
	// state := randomString(64)
	// fmt.Println(state)
	// session.Values["state"] = state
	// session.Save(r, w)
	session, err := gothic.Store.Get(r, gothic.SessionName)
	// if err != nil {
	// 	log.Println(err)
	// 	// send error rest.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	id, _ := session.Values["userID"].(string)
	// if !ok {
	// 	// send error
	// 	return
	// }
	log.Println(session.Values["userID"])
	// session.Values["userID"] = user.ID
	// err = session.Save(req.Request, w.(http.ResponseWriter))
	// if err != nil {
	// 	rest.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	var data = struct{ ID string }{id}
	//{apiConfig.Google.ClientID, state}

	err = indexTemplate.Execute(w, data)
	if err != nil {
		log.Println("error rendering template:", err)
		return
	}
}

func GetUserID(req *rest.Request) (string, error) {
	session, err := gothic.Store.Get(req.Request, gothic.SessionName)
	if err != nil {
		return "", err
	}
	userID, ok := session.Values["userID"].(string)
	if !ok {
		return "", nil
	}
	return userID, nil
}
