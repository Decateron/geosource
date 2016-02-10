package main

import (
	"./fields"
	"encoding/json"
	"errors"
)

type Channel struct {
	// Name of the channel.
	Name string `json:"name"`
	// Fields is usually of type []*Field, but can be if type []json.RawMessage
	// to allow for efficient unparsing.
	Fields interface{} `json:"fields"`
}

// Unmarshals the given JSON blob, returning a Channel on success, or an error
// if unsuccessful. All fields must be valid and empty for parsing to succeed.
func UnmarshalChannel(blob []byte) (*Channel, error) {
	var jsonFields []json.RawMessage
	channel := Channel{
		Fields: &jsonFields,
	}
	json.Unmarshal(blob, &channel)

	channelFields := make([]*fields.Field, len(jsonFields))

	for i, jsonField := range jsonFields {
		field, err := fields.UnmarshalField(jsonField)
		if err != nil {
			return nil, err
		}
		if !field.IsEmpty() {
			return nil, errors.New("Non-empty field submitted")
		}
		channelFields[i] = field
	}
	channel.Fields = channelFields

	return &channel, nil
}

// Unmarshals the given JSON blob into a Submission, and attempts to validate
// and return it in Post form. If the submission is invalid due to either an
// unmarshalling error or a form mismatch, an error is returned.
func (channel *Channel) UnmarshalSubmission(blob []byte) (*Post, error) {

	channelFields, ok := channel.Fields.([]*fields.Field)
	if !ok {
		return nil, errors.New("Invalid type of Fields in channel.")
	}

	var jsonValues []json.RawMessage
	submission := Submission{
		Values: &jsonValues,
	}
	json.Unmarshal(blob, &submission)
	if len(jsonValues) != len(channelFields) {
		return nil, errors.New("An invalid number of values were provided.")
	}

	post := Post{
		Title:   submission.Title,
		Channel: submission.Channel,
		Fields:  make([]*fields.Field, len(channelFields)),
	}

	for i, field := range channelFields {
		fieldForm, ok := field.Form.(fields.Form)
		if !ok {
			return nil, errors.New("Invalid type of Form in field")
		}
		value, err := fieldForm.UnmarshalValue(jsonValues[i])
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
