package api

import (
	"../config"
	"errors"
	"fmt"
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/gplus"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"html/template"
	"log"
	"net/http"
	"os"
)

var apiConfig *config.Config

var ErrInsufficientPermission error = errors.New("Insufficient permission.")
var indexTemplate = template.Must(template.ParseFiles("../app/index.html"))

func Init(config *config.Config) {
	apiConfig = config
	oauthConfig = &oauth2.Config{
		ClientID:     config.Google.ClientId,
		ClientSecret: config.Google.ClientSecret,
		// Scope determines which API calls you are authorized to make
		Scopes:   []string{"https://www.googleapis.com/auth/plus.login", "https://www.googleapis.com/auth/userinfo.email"},
		Endpoint: google.Endpoint,
		// Use "postmessage" for the code-flow for server side apps
		RedirectURL: "postmessage",
	}
	goth.UseProviders(
		gplus.New(config.Google.ClientId, config.Google.ClientSecret, config.Google.CallbackUrl),
		// facebook.New(config.Facebook.ClientId, config.Facebook.ClientSecret, config.Facebook.Callback),
		// twitter.New(config.Twitter.ClientId, config.Twitter.ClientSecret, config.Twitter.Callback),
	)
}

func MakeHandler() (http.Handler, error) {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		// Login
		rest.Post("/login", Login),
		rest.Post("/logout", Logout),

		// Auth
		rest.Get("/auth/#provider", BeginAuth),
		rest.Get("/auth/#provider/callback", CallbackAuth),

		// Channels
		rest.Get("/channels", GetChannels),
		rest.Get("/channels/#channelname", GetChannel),
		// rest.Put("/channels/#channelname", SetChannel),
		rest.Post("/channels", AddChannel),
		rest.Delete("/channels/#channelname", RemoveChannel),

		// Moderators
		rest.Get("/channels/#channelname/moderators", GetModerators),
		rest.Post("/channels/#channelname/moderators", AddModerator),
		rest.Delete("/channels/#channelname/moderators/#username", RemoveModerator),

		// Viewers
		rest.Get("/channels/#channelname/viewers", GetViewers),
		rest.Post("/channels/#channelname/viewers", AddViewer),
		rest.Delete("/channels/#channelname/viewers/#username", RemoveViewer),

		// Banned
		rest.Get("/channels/#channelname/bans", GetBans),
		rest.Post("/channels/#channelname/bans", AddBan),
		rest.Delete("/channels/#channelname/bans/#username", RemoveBan),

		// Posts
		rest.Get("/posts", GetPosts),
		rest.Get("/posts/#pid", GetPost),
		// rest.Put("/posts/#pid", SetPost),
		rest.Post("/posts", AddPost),
		rest.Delete("/posts/#pid", RemovePost),

		// Comments
		rest.Get("/posts/#pid/comments", GetComments),
		rest.Get("/posts/#pid/comments/#cid", GetComment),
		// rest.Put("/posts/#pid/comments/#cid", SetComment),
		rest.Post("/posts/#pid/comments", AddComment),
		rest.Delete("/posts/#pid/comments/#cid", RemoveComment),

		// Users
		rest.Get("/users", GetUsers),
		rest.Get("/users/#username", GetUser),
		// rest.Put("/users/#username", SetUser),
		rest.Post("/users/#username", AddUser),
		rest.Delete("/users/#username", RemoveUser),

		// Subscriptions
		rest.Get("/users/#username/subscriptions", GetSubscriptions),
		rest.Post("/users/#username/subscriptions", AddSubscription),
		rest.Delete("/users/#username/subscriptions/#channelname", RemoveSubscription),
	)
	if err != nil {
		return nil, err
	}
	api.SetApp(router)
	return api.MakeHandler(), nil
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, SESSION_NAME)
	state := randomString(64)
	fmt.Println(state)
	session.Values["state"] = state
	session.Save(r, w)

	var data = struct {
		ClientId, State string
	}{apiConfig.Google.ClientId, state}

	err := indexTemplate.Execute(w, data)
	if err != nil {
		log.Println("error rendering template:", err)
		return
	}
}
