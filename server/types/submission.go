package types

import (
	"github.com/joshheinrichs/geosource/server/types/fields"
)

// A submission contains the necessary information to construct a Post.
type Submission struct {
	Title    string         `json:"title"`
	Channel  string         `json:"channel"`
	Location Location       `json:"location"`
	Values   []fields.Value `json:"values"`
}
