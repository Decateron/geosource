# api

This document outlines the various requests that can be made to the API.

### Auth
Method | URL | Description
-------|-----|------------
GET    | `/auth/#provider` | Calls off to the provider's OAuth2 API to sign them in. Only Google is supported currently, although support could easily be added for many other platforms including Facebook and Twitter.
GET    | `/auth/#provider/callback` | 
POST   | `/logout` | Logs the user out of the application, deleting their session.

### Channels

Method | URL | Description
-------|-----|------------
GET    | `/channels` |
GET    | `/channels/#channelname` |
POST   | `/channels` |
DELETE | `/channels/#channelname` |

### Posts

Method | URL | Description
-------|-----|------------
GET    | `/posts` | 
GET    | `/posts/#postID` | 
POST   | `/posts` |
DELETE | `/posts/#postID` |

### Comments

Method | URL | Description
-------|-----|------------
GET    | `/posts/#postID/comments` |
POST   | `/posts/#postID/comments` |
DELETE | `/posts/#postID/comments/#commentID` |

### Favorite

Method | URL | Description
-------|-----|------------
GET    | `/users/#userID/favorites` |
POST   | `/users/#userID/favorites` |
DELETE | `/users/#userID/favorites/#postID` |

### Subscriptions

Method | URL | Description
-------|-----|------------
GET    | `/users/#userID/subscriptions` |
POST   | `/users/#userID/subscriptions` |
DELETE | `/users/#userID/subscriptions/#channelname` |

