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
	Title   string `json:"title"`
	Channel string `json:"channel"`
	// Values is usually of type []*Value, but can be if type []json.RawMessage
	// to allow for efficient unparsing.
	Values interface{} `json:"values"`
}
