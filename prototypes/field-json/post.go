package main

import (
	"./fields"
)

type Post struct {
	Title   string          `json:"title"`
	Channel string          `json:"channel"`
	Fields  []*fields.Field `json:"fields"`
}

type Submission struct {
	Title   string         `json:"title"`
	Channel string         `json:"channel"`
	Values  []fields.Value `json:"values"`
}
