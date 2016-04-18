# Design

### Website

The website for this project was built in [Polymer 1.0](https://www.polymer-project.org/1.0/). Polymer is based around the idea of web components -- building websites out of small, reusable pieces through the use of custom HMTL tags and HTML imports.

Organization of the files was based on the approach used by the [Google IO 2015](https://github.com/GoogleChrome/ioweb2015) project, built in Polymer 0.5, where custom elements are placed inside the `elements/` directory, and all external web components and dependencies which are used are placed inside the `bower_components/` directory. After the website already under development, other projects such as [polymer-starter-kit](https://github.com/PolymerElements/polymer-starter-kit) have popped up, which may provide slightly better practices for project structuring.

- HTML5 location
- Service workers for caching website via platinum elements
- IndexedDB for storing away posts
	- Alternatives such as local storage have very limited sizes
	- Images stored as base64 strings as there was not support for binary blobs on all browsers when developed

- HTML5 forms and validation not used due to `iron-form` not supporting image inputs at the time: https://github.com/PolymerElements/iron-form/issues/54
	- some support has been added and so this could be 'fixed', but this would require redesign of channel and post creation to handle the change in how data is delivered and as such has not been added
	- as well, the value of a form file input cannot be modified programatically, so loading a saved file off of a user's harddrive or out of the website's storage is not possible
	- html5 forms also limited in terms of types
		- could not support geolocation as a field type
	- less flexible than my structure

### Server

The server for this project was built in [Go 1.5](https://golang.org/). There are several tools that I personally used, which I would suggest adding to your toolbelt when working on the server:

* [**gofmt**](https://golang.org/cmd/gofmt/) - Enforces standard styling conventions to Go code. Using a text editor plugin such as [GoSublime](https://github.com/DisposaBoy/GoSublime) will run this automatically every time you save a file.
* [**golint**](https://github.com/golang/lint) - Points out style mistakes such as undocumented public functions.
* [**go vet**](https://golang.org/cmd/vet/) - Reports likely errors such as calls to Printf with incorrect argument types.
* [**goimports**](https://godoc.org/golang.org/x/tools/cmd/goimports) - Automatically fixes some missing or unnecessary dependencies. This too can be set to run whenever you save a Go file in [Sublime](http://michaelwhatcott.com/gosublime-goimports/).

The server serves all of the static files inside the `app/public/` directory and exposes a RESTful API through which all interactions occur. The website is only available via HTTPS (HTTP simply redirects to HTTPS). This is not only more secure, but faster as well with HTTP/2, which is supported as of Go 1.6 and will be default in Go 1.7.

The server is split into three main packages, `api`, `types`, and `transactions`. The `api` package contains the logic for the RESTful API calls that can be performed. the `types` package contains all of the server-side representations of the application's datatypes such Posts and Channels. It also contains within it the `fields` package, which holds all of the various field types that can used, such as Images and Text. The `transactions` package contains all of the various interactions between the server and the database.

To get an idea of how the packages generally work together, I've described how posts submissions are handled on the server below:

1. A POST request is made to `/api/posts`
2. The request body is dynamically parsed into a Post struct 
	* This checks that correct types were given for the information
3. That information is validated, both checking and potentially modifying the information
	* This both checks and potentially makes the provided information legal
	* Leading and trailing whitespace is removed from most strings
	* Channel names are verified to be exclusively alphanumeric characters 
	* Images are converted to jpgs and saved to stable storage
	* Required post fields are checked to be filled out
4. If the information is deemed valid, the time at which the post was created and the ID of the user is assigned to the post
5. The post it is submitted to the database, along with the requesting user's ID
	* This checks step that the user has permission to submit the information
6. If everything is successful, a response is sent to the user which contains the created post so that it can be added to the user's page without refreshing

- test.sh
- config file to keep track of setup
	- some people use bash variables, harder to keep track of imo

##### Fields

When adding a new field type to the server, a few steps must be taken. First, a constant identifier should be created for the type. This allows the type of the type to be easily identified and handled on both the website and the server. If you were creating a number type, you'd want to add the constant `TypeNumber = "number"` to `fields.go`. Then, you would want to create a new file called `number.go`, which would contain two types: `NumberForm`, which implements the `Form` interface, and `NumberValue`, which implements the `Number` interface. These two types are necessarily highly coupled, and allow for a large amount of code reuse, as the `Form` both validates and understands how to unparse the JSON representation of its associated `Value`. Finally, you would want to add `TypeNumber` as a case to `UnmarshalForm()` function within `fields.go`. 

### Database

The database for this project was built in [PostgreSQL 9.4](http://www.postgresql.org/docs/9.4/static/release-9-4.html), as it has support for some NoSQL features such as JSON, which were used within this project to store dynamic content. While this could've been modeled relationally, it would ultimately be significantly more complex and less efficient.

- general design
	- static content stored traditionally
	- dynamic content such as channel forms, post contents stored dynamically as JSON
	- PostGIS used for spatial indexing and queries

##### Users

- username currently just set to whatever name given by OAuth providers
	- couldn't find limit on length (at least for google), so right now is just set to TEXT, ideally this would be changed at some point though, as realistically a limit should be set
	- avatar is the URL of the user's avatar from the provider
	- userID is base64 encoded UUID
		- somewhat random user ID, mainly an attempt to avoid using serial identifiers

##### Channels

- channelnames currently limited to alphanumeric characters, no spaces
	- could easily add support for them
	- either give chanenls base64uuid ID or URL encode channelname which is currently used as identifier within rest api
		- former probably better since spaces should not be used in rest API
	- visibility not currently in use
	- field values should be empty
