# Design

The overall design of the GeoSource follows the standard 3 layer architecture approach, with a website, server, and database. These three moduals are mostly seperate, handling the logic of their respective domains. However, the server does limit what sorts of information can be inserted into the database beyond what is specified within SQL. While this probably could've been handled via SQL scripts, I thought that putting the logic within the server would make it easier to move to different datastores in the future.

## Website

The website for this project was built in [Polymer 1.0](https://www.polymer-project.org/1.0/). Polymer is based around the idea of web components -- building websites out of small, reusable pieces through the use of custom HMTL tags and HTML imports.

Organization of the files was based on the approach used by the [Google IO 2015](https://github.com/GoogleChrome/ioweb2015) project, built in Polymer 0.5, where custom elements are placed inside the `elements/` directory, and all external web components and dependencies which are used are placed inside the `bower_components/` directory. After the website was already under development, other projects such as [polymer-starter-kit](https://github.com/PolymerElements/polymer-starter-kit) have popped up, which may provide slightly better practices for project structuring.

### Required Technologies

This project relies on a large number of recent web features. While most are not needed to browse the website, many are required for either posting or accessing the website while offline. These technologies include:

- **HTML5 Geolocation** - Needed to identify the location at which the user is recording a post.
- **IndexedDB** - Needed for saving posts. I attempted to use other more widely adopted technologies such as local storage, but they were serverly limited in size (only able to store a few MB of data on some devices), making it impossible to store posts with multiple images. Files have to be stored within the website as there is no way to reference an external file for later use for security reasons. As a result the data has to be duplicated, or else the user would have to rereference the files when submitting at a later date. Neither option is great, but duplicated data seems like the better of the two at least in the case of images.
- **Service Workers** - This is needed for accessing the website while offline. Service workers can be associated with a webpage and essentially.
- **Web Components** - This technically isn't needed, as there is a [polyfill](https://en.wikipedia.org/wiki/Polyfill) for it, but polyfills aren't cheap. 

## Server

The server for this project was built in [Go 1.6](https://golang.org/). There are several tools that I personally used, which I would suggest adding to your toolbelt when working on the server:

* [**gofmt**](https://golang.org/cmd/gofmt/) - Enforces standard styling conventions to Go code. Using a text editor plugin such as [GoSublime](https://github.com/DisposaBoy/GoSublime) will run this automatically every time you save a file.
* [**golint**](https://github.com/golang/lint) - Points out style mistakes such as undocumented public functions.
* [**go vet**](https://golang.org/cmd/vet/) - Reports likely errors such as calls to Printf with incorrect argument types.
* [**goimports**](https://godoc.org/golang.org/x/tools/cmd/goimports) - Automatically fixes some missing or unnecessary dependencies. This too can be set to run whenever you save a Go file in [Sublime](http://michaelwhatcott.com/gosublime-goimports/).

The server serves all of the static files inside the `app/public/` directory and exposes a RESTful API through which all interactions occur. The website is only available via HTTPS (HTTP simply redirects to HTTPS). This is not only more secure, but faster as well with HTTP/2, which is supported by default as of Go 1.6 and some webcomponent minification tools such as  [vulcanize](https://github.com/Polymer/vulcanize) obselete.

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

### Fields

When adding a new field type to the server, a few steps must be taken. First, a constant identifier should be created for the type. This allows the type of the type to be easily identified and handled on both the website and the server. If you were creating a number type, you'd want to add the constant `TypeNumber = "number"` to `fields.go`. Then, you would want to create a new file called `number.go`, which would contain two types: `NumberForm`, which implements the `Form` interface, and `NumberValue`, which implements the `Number` interface. These two types are necessarily highly coupled, and allow for a large amount of code reuse, as the `Form` both validates and understands how to unparse the JSON representation of its associated `Value`. Finally, you would want to add `TypeNumber` as a case to `UnmarshalForm()` function within `fields.go`. 

### Testing

To run the tests for the server, all you need to do is be located within the `server/` directory and run:

```
go test ./...
```

This will recursively tests all of the folders within the `server/` folder. A `test.sh` file has also been included which generates proper code coverage reports. Test coverage on the server is somewhat minimal, the approach to testing each of the packages has at least been grappled with, with clean databases been created on Travis CI to test the `transactions` package, testing API calls in the `api` package, and standard unit tests of the `types` package.

## Database

The database for this project was built in [PostgreSQL 9.4](http://www.postgresql.org/docs/9.4/static/release-9-4.html), as it has support for some NoSQL features such as JSON, which was used within this project to store dynamic content. While this could've been modeled relationally, it would ultimately be significantly more complex and less efficient.

![](https://joshheinrichs.github.io/geosource/er-diagram.png)

An ER diagram of the database has been included above. I think the structure is generally self-explanitory, so I won't go into much detail about it. The one thing I will mention is that I think the three permission tables, `channel_viewers`, `channel_bans`, and `channel_moderators` could be combined into a single `channel_permissions` table, but these tables aren't currently being used, so I haven't revised the structure.

### Users

- **p_userid** - Users are given a base64 encoded UUID when they are created. These are somewhat random, although I mainly wanted to avoid using serial identifiers.
- **p_email** - From my research, it appears that emails cannot be longer than 254 characters, so I limit it to strings of that length.
- **p_username** - Currently the username is just set to whatever is given by the OAuth providers. Google doesn't seem to have any limit. on length, so I just set it to TEXT. A limit should probably be set at some point in the future.
- **p_avatar** - The maximum length of a URL is apparently 2083 characters, so I set the type to TEXT for simplicity's sake.

### Channels

- **ch_channelname** - The name of the channel is currently limited to 20 characters, and must be unique. The length is somewhat arbitrary, and really could be extended if needed. Currently, channels can only consist of alphanumeric characters, but support for other characters could be added by either URL encoding the name when referencing it via API urls, or giving each channel an associated UUID identifier, similar to posts and users.
- **ch_userid_creator** - The user which created the channel.
- **ch_visibility** - Can be set to `'public'` or `'private'`. Currently recorded but not taken into account when using the site.
- **ch_fields** - The channel's associated form structure, encoded in JSON.

### Posts

- **p_postid** - Posts are given a base64 encoded UUID when they are created. These are somewhat random. I mainly wanted to prevent people from being able to linearlly access every post through API requests and know how many private posts are in the database if support for private channels is added.
- **p_userid_creator** - Creator of the post.
- **p_channelname** - The channel to which this post was posted.
- **p_title** - The title of the post, currently limited to 140 characters. This is somewhat arbitrary and could be changed. 
- **p_thumbnail** - This field stores the URL of the post's associated thumbnail. The URLs are relative to the website, so only `/media/images/uuid.jpg` is stored. In retrospect, it may be good to change this to TEXT and allow for any URL, so that moving static. content to a CDN or something would be possible without restructuring the database. 
- **p_time** - The time at which the post was submitted.
- **p_location** - The location of the post. This makes use of PostGIS's geometry library to allow for spatial indexing.
- **p_fields** - The post's associated form and value, encoded in JSON. Both the form and values are stored for each post  so that the channel's form could change in the future without breaking posts which have already been submitted to the channel.
