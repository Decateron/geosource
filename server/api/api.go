package api

import (
	"errors"
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
)

var ErrInsufficientPermission error = errors.New("Insufficient permission.")

func MakeHandler() (http.Handler, error) {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		// Login
		// rest.Post("/login", Login),
		// rest.Get("/logout", Logout),

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
