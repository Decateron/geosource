package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/joshheinrichs/geosource/server/types/fields"
)

const (
	MAX_CHANNELNAME_LEN = 20
)

var channelNameRegexp = regexp.MustCompile("^\\w+$")

type Channel struct {
	Name       string        `json:"name" gorm:"column:ch_channelname"`
	CreatorID  string        `json:"creatorID" gorm:"column:ch_userid_creator"`
	Visibility string        `json:"visibility" gorm:"column:ch_visibility"`
	Fields     fields.Fields `json:"fields" gorm:"column:ch_fields" sql:"type:JSONB NOT NULL"`
}

// Unmarshals the given JSON blob, returning a Channel on success, or an error
// if unsuccessful. All fields must be valid and empty for parsing to succeed.
func UnmarshalChannel(blob []byte) (*Channel, error) {
	unmarshalChannel := struct {
		Channel
		JsonFields []json.RawMessage `json:"fields"`
	}{}
	json.Unmarshal(blob, &unmarshalChannel)

	channelFields := make([]*fields.Field, len(unmarshalChannel.JsonFields))
	for i, jsonField := range unmarshalChannel.JsonFields {
		field, err := fields.UnmarshalField(jsonField)
		if err != nil {
			return nil, err
		}
		if !field.IsEmpty() {
			return nil, errors.New("Non-empty field submitted")
		}
		channelFields[i] = field
	}
	unmarshalChannel.Channel.Fields = channelFields

	return &unmarshalChannel.Channel, nil
}

// Validates the channel. Returns an error if any fields are invalid, or nil
// otherwise.
func (channel *Channel) Validate() error {
	channel.Name = strings.TrimSpace(channel.Name)
	if len(channel.Name) == 0 {
		return errors.New("Channel name cannot be empty.")
	} else if len(channel.Name) > MAX_CHANNELNAME_LEN {
		return errors.New(fmt.Sprintf("Length of channel name cannot exceed %i characters.", MAX_CHANNELNAME_LEN))
	} else if !channelNameRegexp.MatchString(channel.Name) {
		return errors.New("Channel names may only contain alpha numeric characters, hyphens or underscores.")
	}
	return nil
}

// Unmarshals the given JSON blob into a Submission, and attempts to validate
// and return it in Post form. If the submission is invalid due to either an
// unmarshalling error or a form mismatch, an error is returned.
func (channel *Channel) UnmarshalSubmissionToPost(blob []byte) (*Post, error) {
	unmarshalSubmission := struct {
		Submission
		JsonValues []json.RawMessage `json:"values"`
	}{}

	json.Unmarshal(blob, &unmarshalSubmission)
	if len(unmarshalSubmission.JsonValues) != len(channel.Fields) {
		return nil, errors.New("An invalid number of values were provided.")
	}

	post := Post{
		PostInfo: PostInfo{
			Title:    unmarshalSubmission.Title,
			Channel:  unmarshalSubmission.Channel,
			Location: unmarshalSubmission.Location,
		},
		Fields: make([]*fields.Field, len(channel.Fields)),
	}

	for i, field := range channel.Fields {
		value, err := field.Form.UnmarshalValue(unmarshalSubmission.JsonValues[i])
		if err != nil {
			return nil, err
		}
		post.Fields[i] = &fields.Field{
			Label:    field.Label,
			Type:     field.Type,
			Required: field.Required,
			Form:     field.Form,
			Value:    value,
		}
		err = post.Fields[i].Validate()
		if err != nil {
			return nil, err
		}
	}
	err := post.GenerateThumbnail()
	if err != nil {
		return nil, err
	}

	return &post, nil
}

// Returns the name of the channel's corresponding table in the database.
func (channel *Channel) TableName() string {
	return "channels"
}
