package field

import (
	"errors"
)

const (
	TYPE_TEXT         = "text"
	TYPE_NUMBER       = "number"
	TYPE_CHECKBOXES   = "checkboxes"
	TYPE_RADIOBUTTONS = "radiobuttons"
	TYPE_IMAGES       = "images"
)

var ErrMissingType error = errors.New("Missing type.")
var ErrInvalidType error = errors.New("Invalid type.")
var ErrMissingLabel error = errors.New("Missing label.")
var ErrInvalidLabel error = errors.New("Invalid label.")
var ErrMissingValue error = errors.New("Missing value.")
var ErrInvalidValue error = errors.New("Invalid value.")

type Field struct {
	Type         string        `json:"type"`
	Label        string        `json:"label"`
	Required     bool          `json:"required"`
	Text         *Text         `json:"text,omitempty"`
	Number       *Number       `json:"number,omitempty"`
	Checkboxes   *Checkboxes   `json:"checkboxes,omitempty"`
	Radiobuttons *Radiobuttons `json:"radiobuttons,omitempty"`
	Images       *Images       `json:"images,omitempty"`
}

type Value interface {
	Validate() error
	IsEmpty() bool
}

func (field *Field) Validate() error {
	// ensure that only one field is filled out
	if field.Type != TYPE_TEXT && field.Text != nil ||
		field.Type != TYPE_NUMBER && field.Number != nil ||
		field.Type != TYPE_CHECKBOXES && field.Checkboxes != nil ||
		field.Type != TYPE_RADIOBUTTONS && field.Radiobuttons != nil ||
		field.Type != TYPE_IMAGES && field.Images != nil {
		return ErrInvalidValue
	}

	var value Value
	switch field.Type {
	case TYPE_TEXT:
		value = field.Text
	case TYPE_NUMBER:
		value = field.Number
	case TYPE_CHECKBOXES:
		value = field.Checkboxes
	case TYPE_RADIOBUTTONS:
		value = field.Radiobuttons
	case TYPE_IMAGES:
		value = field.Images
	default:
		return ErrInvalidType
	}

	err := value.Validate()
	if err != nil {
		return err
	}
	if field.Required && value.IsEmpty() {
		return ErrInvalidValue
	}
	return nil
}
