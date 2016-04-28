# Reflection

I feel like I gained a good amount of insight into the various technologies I used

### What Worked

##### Go
- Lots of libraries for RESTful APIs, SQL transactions, etc.
- Type system is nice for validation
	- Very certain about what does and doesn't make it into database
	- Makes refactoring less scary
- Great support for testing

##### Travis CI
- minimal work to support wide array of tests
	- image conversion
	- database transactions
	- standard unit tests

##### PostgreSQL
- NoSQL features such as JSON type made for nice combination of static and dynamic content


##### Prototyping
- Verified technologies individually, easy to hook them togethor


### What Didn't Work

##### Polymer
- might be fine, I found it frustrating
	- poor documentation in general for supported and unsupported features
	- difficult to work with large datastructures

##### PostgreSQL at a larger scale
- complex queries prove difficult to optimize
- basically linear scaling which is about as bad as it gets

##### User Interaction
- doesn't feel very simple or intuitive
	- no good examples of spatiotemporal browsing to go off of
	- could be an interesting project for someone

### Problems and Solutions

One thing I noticed was that information was stored with various amounts of detail depending upon the state of the program. I 

- custom forms
- storing away files

### Future Work

- public and private channels

- moderation tools
- permission levels
	- banned, viewers, moderators, owner and admin permission checked for various transactions
	- make single permission table for channels rather than 3
		- doesn't make sense for someone to be a moderator and banned
		- just make an enum or something for permission types
		- channelname, username, permission

- investigate more HTML5 standard form submissions and validation
	- these are a bit of a pain due to limited naming conventions
	- have to have a standard string naming convention thats understood
	- have to deal with files differently than standard

- general improvements to error handling
	- better HTTP error codes and messages on server side
	- better display of errors on website

- try NoSQL solution for spatiotemporal querying
	- O(1) access time can be achieved for 20 most recent posts within area with support for lazy loading!!!
		- will require some work to implement, limitations on querying, and a good amount of data redundancy
		- based on screen region, approximate to a bunch of squares
			- try to keep the number of squares somewhat small
			- the larger the screen size, the larger squares can be used
			- each square contains a list of posts ordered by recency
			- to get the first 20 posts from a set of lists, need to grab the first 20 posts from each list, merge until 20 most recent are found
				- this is the main performance bottleneck of the algorithm, but can technically be constant
				- keep track of how deep you've gotten into each list, and the size of each list and send that info over to the user
					- this will provide sufficient information to know where to start for each of these squares when attempting to load the next 20 posts

- optimizing website with minification and stuff
	- honestly I'd rather it was rebuilt from ground up in something other than polymer
		- google uses it in a few of their sites but i'm not convinced personally

- routing - allow for linking of posts
	- i.e. geosource.com/posts/#postID opens that post on the website

- audio and video support
	- not sure what the library support is like in Go, ffpmpeg bindings available but no simple libraries at this time
	- check Go-Awesome, may have some up to date suggestions

- general usability investigation
	- not much thought was put into usability of interface
	- how best to support spatial-temporal browsing?
		- doesn't feel very natural at the moment

- better testing and documentation
- automated deployment
	- preferably with coverage, syntax, documentation requirements

- investigate into nginx as a proxy for go server
	- remove need for root access to run on port 80 and 443

- live updates using websockets
	- lots of great concurrency primitives in Go to support this

- take advantage of EXIF data when converting images
