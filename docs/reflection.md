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

![](https://joshheinrichs.github.io/geosource/database-benchmark.png)

##### User Interaction
- doesn't feel very simple or intuitive
	- no good examples of spatiotemporal browsing to go off of
	- could be an interesting project for someone

### Problems and Solutions

One thing I noticed was that information was stored with various amounts of detail depending upon the state of the program. I 

- custom forms
- storing away files
- HTML5 forms and validation not used due to `iron-form` not supporting image inputs at the time: https://github.com/PolymerElements/iron-form/issues/54
	- some support has been added and so this could be 'fixed', but this would require redesign of channel and post creation to handle the change in how data is delivered and as such has not been added
	- as well, the value of a form file input cannot be modified programatically, so loading a saved file off of a user's harddrive or out of the website's storage is not possible
	- html5 forms also limited in terms of types
		- could not support geolocation as a field type
	- less flexible than my structure

### Future Work

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
- Optimizing website with minification and stuff
	- Honestly I'd rather it was rebuilt from ground up in something other than Polymer
		- I'd suggest looking into React and Redux 
		- Google uses Polymer it in a few of their sites but personally I'm not convinced
- Routing - allow for direct links to posts
	- i.e. geosource.com/posts/#postID opens that post on the website

- Audio and video support
	- Not sure what the library support is like in Go, ffpmpeg bindings available but no simple libraries at this time
	- Check Go-Awesome, may have some up to date suggestions
	- Will also have to worry about how that information is saved away on the website while offline

- General usability improvements
	- Not much thought was put into usability of interface
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
