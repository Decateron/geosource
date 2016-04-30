# Reflection

Overall, I think I gained a good amount of inisght into the strengths and weaknesses of the various technologies I used.

## What Worked

### Go

Overall, I found that Go worked very well for developing a robust server. The type system was nice for providing sanity checks and ensuring that no invalid information gets into the database without really getting in the way. The way error checking is generally handled also makes handling errors within requests very clear and explicit, as whenever you call a function where an error could potentially occur, you just check if an error occured and handle it explicity. There are also lots of great packages for building RESTful APIs and SQL queries, converting images, interacting with OAuth APIs, and more. Go also has built in support for testing, benchmarking and generating code coverage reports.

### Travis CI

I thought Travis CI worked very well for continuous integration. It removed most of the headache associated with setting up a Jenkins server since it sets up a clean VM each time it runs the set of tests, and comes with many of the dependencies that I needed already installed such as Go, PostgreSQL and ImageMagick. The documentation isn't perfect however, and it was sometimes difficult to figure out what versions of programs came with different Linux distrubtions that were available.

### PostgreSQL

Desipite its name, PostgreSQL has support for a wide array of NoSQL features. Although there were many NoSQL features which I didn't use (such as PostgreSQL's key value store capabilities), the one feature that I did use, JSON, was invaluable for consisely modeling the database since dynamic content doesn't map well to relational modeling. Overall, I thought this made for a nice mix of the two paradigms when working with small amounts of data.

PostgreSQL also has support for spatial querying via PostGIS, although this didn't end up being very useful at least in the case of query performance, as is documented below. However, PostGIS can enable much more complex spatial querying than what I ended up using it for, so it might be interesting utilizing it more heavily in the future to say find all the posts within Saskatoon, rather than just the user's screen region.

### Prototyping

I think that prototyping was very helpful for figuring out what technologies worked and didn't work, and helped to reduce the number of risks in building this project. It would've been painful learn that some browser features on which this application relied didn't work as expected, and prevented the completion of the project. Thankfully this wasn't the case, but it was still useful for identifying major issues earlier into development, such as: 

- HTML5 Geolocation functionality while offline
	- Works fine for devices with built in GPS, but not devices such as laptops which rely on the user's IP to estimate their location
-  File storage limitations
	- Local Storage is extremely limited, so I ended up having to use IndexedDB instead
	- IndexedDB does not allow for binary blob storage, so I had to convert to the data to base64 to save it

## What Didn't Work

### Polymer

While Polymer is apparently pretty good and has been used in websites at Google including Google Play Music, I personally found it frustrating to deal with, largely due to its infancy. First, its documentation of limitiations is rather sparse, so I often found myself running into issues where I would have to search through several GitHub issues and StackOverflow pages before I found it was some limitation due to how Polymer was implemented. Many of the official libraries also contained undocumented limitaitons which could only be noticed by either filing an issue on GitHub, or digging through the source code, both of which are not ideal. Due to these limitations as well as some bugs which I encountered, I found myself somewhat frequently having to build around Polymer rather than with it.

### PostgreSQL at a Large Scale

Although PostgreSQL with JSON made it decently simple to model a functioning database, it doesn't scale very well to large amounts of data, even with the spatial indexing provided by PostGIS. Ideally when browsing you'd probably want to load the first 20 or so posts from the user's current screen region, and then lazy load additional posts as they scroll through the list on the size. The problem with this approach which is pretty standard in social networks such as Twitter and Facebook is that because data has to be filtered both by its location and by its recency, there aren't any good indexing methods out there that can provide this type of query efficiently. From the simple benchmarks I ran, trying to find the 20 most recent posts within a region from a set of posts at random locations, it seems like the query time scales pretty much linearlly, which is about as bad as one could expect.

![](https://joshheinrichs.github.io/geosource/database-benchmark.png)

There are definetly ways to improve this by limiting the flexibility of queries. I think this tradeoff should be decided upon after understanding this application a bit better from a usability perpsective since ultimately that's more important than if the application will work smoothly with millions of posts.

### User Interaction

User interaction with spatio-temporal browsing is still somewhat of a grey-area to me. While I can personally navigate through the system pretty well, I don't think it would feel very natural to most people. I think a lot of work could still be done in this area, as I feel it's one of the largest weaknesses of the current system.

### Go Dependancy Management

Up until Go 1.6 which was released relatively recently, Go did not have an agreed upon method for keeping track of the version of a package on which your code depenended. I ran into a few cases during development where my tests on Travis CI would break due to updates to some external packages which I used. Since there was no way to specify that I was using an older version of the package, I'd have to spend a few hours updating my server. In Go 1.6, support for vendoring was added which addresses this issue, although I haven't had time to investigate adding it to the project.

## Problems and Solutions

### Saving Posts

Saving away posts in a website has some challenges associated with it. The most notable problem is that websites don't have direct access to the filesystem for security reasons. As such, if you want to keep a photo that the user's taken associated with a post so that they don't have to reselect all the files before uploading, you have to save it along with the rest of the post inside of IndexedDB, which means converting the image into base64 when storing it (since binary blobs are not supported by all browsers). While this is a bit hacky, it works alright in practice, at least for images and audio files. That said, it wouldn't really work for videos, which can be multiple gigabytes in size. Storing videos in base64 would be both computationally expensive and take up a lot of extra storage space on the user's device which is pretty desirable.

### Dynamic Forms

I went through a lot of iterations on the server trying to support dynamic forms cleanly. Initially I was attempting to do a single pass of static JSON unmarshaling, which ended up not being very clean, as a field would have to cotain all of the potential types within the struct, which minimzied code reuse i.e.:

```go
type Field struct {
   Label        string       `json:"label"`
   Text         *string      `json:"text"`
   Radiobuttons Radiobuttons `json:"radiobuttons"`
   Checkboxes   Checkboxes   `json:"checkboxes"`
   Images       Images       `json:"images"`
   ...
}
```

Ultimately I decided on creating two seperate interfaces which a given type would have to implement, a `Form` and a `Value`, and giving all fields the same structure i.e.:

```go
type Field struct {
   Type  string `json:"type"`
   Label string `json:"label"`
   Form  Form   `json:"form"`
   Value Value  `json:"value"`
}
```

So the images type has associated `ImagesForm` and `ImagesValue` structs which implement the `Form` and `Value` interfaces. This actually makes the JSON unmarshaling pretty clean since it is still mostly automatic, although you have to switch over the type string to properly unmarshal the Form. It also makes it pretty easy to logically seperate channel forms from posts, since channel forms shouldn't have an associated value. Splitting up the form and value also makes it easier to deal with user submissions, as they only have to pass in an array of values which need to be validated against the form which is stored in the database, which means that less information has to be sent by the user and validated by the server.

## Future Work

I think there's a lot of ways in which this project could be continued:

- Additional filtering/searching options
	- Textual search
	- Filtering by channel
- Public and private channels
- Moderation tools
- Add permissions
	- Banned, viewers, moderators, owner and admin permission checked for various transactions
	- Make single permission table for channels rather than 3
		- Doesn't make sense for someone to be a moderator and banned
		- Just make an enum or something for permission types
		- Channelname, username, permission
- General improvements to error handling
	- Better HTTP error codes and messages on server side
	- Better display of errors on website
- Try NoSQL solution for spatiotemporal querying 
	- O(1) access time can be achieved with Redis for 20 most recent posts within area with support for lazy loading
		- Will require some work to implement, severly limits types of queries, and requries a good amount of data redundancy
		- Based on screen region, approximate to a bunch of squares
			- Try to keep the number of squares somewhat small
			- Each square contains a list of posts ordered by recency
			- To get the first 20 posts from a set of lists, need to grab the first 20 posts from each list, merge until 20 most recent are found
				- This is the main performance bottleneck of the algorithm, but is technically be constant
				- Keep track of how deep you've gotten into each list, and the size of each list and send that info over to the user
					- This will provide sufficient information to know where to start for each of these squares when attempting to load the next 20 posts
- Optimizing the website with minification and stuff
	- Honestly I'd rather it was rebuilt from ground up in something other than Polymer
		- I'd suggest looking into React and Redux 
- Routing to allow for direct links to posts
	- i.e. geosource.usask.ca/posts/#postID would open directly to the post with ID #postID on the website
- Audio and video support
	- Not sure what the library support is like in Go, ffpmpeg bindings available but no simple libraries at this time
	- Check Go-Awesome, may have some up to date suggestions
	- Will also have to worry about how that information is saved away on the website while offline
- General usability improvements
	- Not much thought was put into usability of the interface
	- How best to support spatial-temporal browsing?
		- Doesn't feel very natural at the moment
- More thorough testing and documentation
- Automated deployment
	- Preferably with coverage, style, and documentation requirements
	- How to handle changes to the database specification?
- Investigate into nginx as a proxy for go server
	- Remove need for root access to run on port 80 and 443
- Live updates using websockets
	- Lots of great concurrency primitives in Go to support this
- Take advantage of EXIF data when converting images
	- Automatically rotate image based on oritentation data etc.
