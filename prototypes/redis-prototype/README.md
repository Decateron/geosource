# api-prototype

A simple prototype for ironing out the basics of the API and the HTML <-> Golang <-> Redis connection.

This is only testing sending and receiving basic information. I don't plan on expanding this prototype anymore as it seems that redis won't be a great fit for some of the more complicated features that have been requested.

### Setup

To run this program, you'll need to have redis and go installed. After that, it's as simple as running:

```
redis-server
go run main.go
```

The site will then be accessible from localhost:8000
