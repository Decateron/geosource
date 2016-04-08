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
	maxChannelnameLength = 20
)

var channelNameRegexp = regexp.MustCompile("^\\w+$")

type ChannelInfo struct {
	Name       string `json:"name" gorm:"column:ch_channelname"`
	CreatorID  string `json:"creatorID" gorm:"column:ch_userid_creator"`
	Visibility string `json:"visibility" gorm:"column:ch_visibility"`
}

// TableName returns the name of ChannelInfo's corresponding table in the
// database.
func (ChannelInfo) TableName() string {
	return "channels"
}

type Channel struct {
	ChannelInfo
	Fields fields.Fields `json:"fields" gorm:"column:ch_fields" sql:"type:JSONB NOT NULL"`
}

// UnmarshalChannel unmarshals the given JSON blob, returning a Channel on
// success, or an error if unsuccessful. All fields must be valid and empty for
// parsing to succeed.
func UnmarshalChannel(blob []byte) (*Channel, error) {
	unmarshalChannel := struct {
		Channel
		JSONFields []json.RawMessage `json:"fields"`
	}{}
	json.Unmarshal(blob, &unmarshalChannel)

	channelFields := make([]*fields.Field, len(unmarshalChannel.JSONFields))
	for i, jsonField := range unmarshalChannel.JSONFields {
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

// Validate validates the channel. Returns an error if any fields are invalid,
// or nil otherwise.
func (channel *Channel) Validate() error {
	channel.Name = strings.TrimSpace(channel.Name)
	if len(channel.Name) == 0 {
		return errors.New("Channel name cannot be empty.")
	} else if len(channel.Name) > maxChannelnameLength {
		return errors.New(fmt.Sprintf("Length of channel name cannot exceed %i characters.", maxChannelnameLength))
	} else if !channelNameRegexp.MatchString(channel.Name) {
		return errors.New("Channel names may only contain alpha numeric characters, hyphens or underscores.")
	}
	err := channel.Fields.ValidateForms()
	if err != nil {
		return err
	}
	return nil
}

// UnmarshalSubmissionToPost unmarshals the given submission in JSON format into
// a post. If the submission is invalid due to either an unmarshalling error or
// a form mismatch, an error is returned.
func (channel *Channel) UnmarshalSubmissionToPost(blob []byte) (*Post, error) {
	unmarshalSubmission := struct {
		Submission
		JSONValues []json.RawMessage `json:"values"`
	}{}

	json.Unmarshal(blob, &unmarshalSubmission)
	if len(unmarshalSubmission.JSONValues) != len(channel.Fields) {
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
		value, err := field.Form.UnmarshalValue(unmarshalSubmission.JSONValues[i])
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
	}
	return &post, nil
}

// TableName returns the name of Channel's corresponding table in the database.
func (Channel) TableName() string {
	return "channels"
}
