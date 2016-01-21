package api

import (
	"code.google.com/p/go-uuid/uuid"
	"encoding/base64"
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/markbates/goth/gothic"
	"net/http"
)

func BeginAuth(w rest.ResponseWriter, req *rest.Request) {
	setProvider(req)
	gothic.BeginAuthHandler(w.(http.ResponseWriter), req.Request)
}

func CallbackAuth(w rest.ResponseWriter, req *rest.Request) {
	setProvider(req)
	_, err := gothic.CompleteUserAuth(w.(http.ResponseWriter), req.Request)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// transaction.GetUser(user.Email)
	// t, err := template.New("foo").Parse(userTemplate)
	// if err != nil {
	// 	rest.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// t.Execute(w.(http.ResponseWriter), user)
	http.Redirect(w.(http.ResponseWriter), req.Request, "https://localhost:8080/", http.StatusTemporaryRedirect)
}

// gothic requires the provider as a query value
func setProvider(req *rest.Request) {
	v := req.Request.URL.Query()
	v.Set("provider", req.PathParam("provider"))
	req.Request.URL.RawQuery = v.Encode()
}

var userTemplate = `<p>Name: {{.Name}}</p>
<p>Email: {{.Email}}</p>
<p>NickName: {{.NickName}}</p>
<p>Location: {{.Location}}</p>
<p>AvatarURL: {{.AvatarURL}} <img src="{{.AvatarURL}}"></p>
<p>Description: {{.Description}}</p>
<p>UserID: {{.UserID}}</p>
<p>AccessToken: {{.AccessToken}}</p>
`
