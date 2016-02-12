package types

import (
	"./fields"
	"time"
)

type Post struct {
	Id        string          `json:"id"`
	CreatorId string          `json:"creator"`
	Channel   string          `json:"channel"`
	Title     string          `json:"title"`
	Time      time.Time       `json:"time"`
	Location  Location        `json:"location"`
	Fields    []*fields.Field `json:"fields"`
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Submission struct {
	Title   string         `json:"title"`
	Channel string         `json:"channel"`
	Values  []fields.Value `json:"values"`
}
