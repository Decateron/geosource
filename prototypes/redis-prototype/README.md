# api-prototype

A simple prototype for ironing out the basics of the API and the HTML <-> Golang <-> Redis connection.

This is only testing sending and receiving basic information, but may be expanded to include google logins, which will be important for authenticating user permissions.

### Setup

To run this program, you'll need to have redis and go installed. After that, it's as simple as running:

```
redis-server
go run main.go
```

The site will then be accessible from localhost:8080
