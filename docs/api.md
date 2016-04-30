# API

This document outlines the various requests that can be made to the API. To see an example response body for a given request, you can usually just append the associated path to the end of the website's URL. In the event of an error, a response will be returned with a HTTP code other than 200 (usually internal server error), and JSON will be returned describing the error. All POST requests respond with the content that was created by the request. 

## Auth
Method | URL | Description
-------|-----|------------
GET    | `/api/auth/#provider` | Calls off to the provider's OAuth2 API to sign them in. Only Google is supported currently, although support could easily be added for many other platforms including Facebook and Twitter.
GET    | `/api/auth/#provider/callback` | The callback URL from OAuth authentication for the given provider.
POST   | `/api/logout` | Logs the user out of the application, deleting their session.

## Channels
Method | URL | Description
-------|-----|------------
GET    | `/api/channels` | Returns the meta information for all channels, including whether the user is subscribed to it.
GET    | `/api/channels/#channelname` | Returns the information for a channel with name #channelname, including its form.
POST   | `/api/channels` | Adds the channel provided that the user is signed in. The request body is expected to be a channel (as defined within the server), with all fields except for `creatorID` provided. The user must be logged in to perform this action.

## Posts
Method | URL | Description
-------|-----|------------
GET    | `/api/posts` | Returns the metainformation about some posts. Posts can be filtered with a query string (the PostQueryParams structure within the server).
GET    | `/api/posts/#postID` | Returns the information for a post with id `#postID`, including its fields.
POST   | `/api/posts` | Adds a post. The request body is expected to be a Submission (as defined within the server). The user must be logged in to perform this action.

## Comments
Method | URL | Description
-------|-----|------------
GET    | `/api/posts/#postID/comments` | Returns the set of comments for the post with ID `#postID`.
POST   | `/api/posts/#postID/comments` | Adds a comment to the post with ID `#postID`. The request body is expected to be a Post (as defined within the server), with all fields except for `id`, `creatorID`, `thumbnail`, and `time` provided. The user must be logged in to perform this action.

## Favorites
Method | URL | Description
-------|-----|------------
POST   | `/api/users/#userID/favorites` | Adds a post to the user's favorites. The request body is expected to contain a field `postID` with the ID of the post to be favorited. This action can only be performed by the user with ID `#userID`.
DELETE | `/api/users/#userID/favorites/#postID` | Removes a post with ID `#postID` from the user's favorites. This action can only be performed by the user with ID `#userID`. 

## Subscriptions
Method | URL | Description
-------|-----|------------
POST   | `/api/users/#userID/subscriptions` | Adds a channel to the user's subscriptions. The request body is expected to contain a field `channelname` with the name of the channel to be subscribed to. This action can only be performed by the user with ID `#userID`.
DELETE | `/api/users/#userID/subscriptions/#channelname` | Removes a channel with name `#channelname` from the user's subscriptions. This action can only be performed by the user with ID `#userID`.

