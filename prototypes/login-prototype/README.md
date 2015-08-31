# login-prototype

Demoing Google+ login. Implementation is based off of the [gplus-quickstart-go](https://github.com/googleplus/gplus-quickstart-go) project, adapted to Polymer. 

This prototype also supports username creation, and stores the data into a simple Postgresql database which records a user's email and username.

A simple, custom signin button was created since the `google-signin` web component does not allow for offline access (since it isn't capable of providing a code). It's a bit buggy at the moment, so hopefully the official `google-signin` component will have that feature included in the future.

### Setup

Follow Step 1 of the [gplus-quickstart-go tutorial](https://developers.google.com/+/web/samples/go) to set up the Google+ API.

After that, you should create a config file named `config.gcfg` based off of the included `example.gcfg` file, which contains your Client ID and Client Secret, and database login information. This config file will be specific to you and should remain private as exposing the Client Secret or database login credentials can create security issues.

After that, it should be as simple as running:

```bash
go run *.go
```

within the `server/` folder. The site will then be accessible via `http://localhost:8000`.
