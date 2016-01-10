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
var ErrMultipleValues error = errors.New("Multiple values.")

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
	value, err := field.GetValue()
	if err != nil {
		return err
	}
	err = value.Validate()
	if err != nil {
		return err
	}
	if field.Required && value.IsEmpty() {
		return ErrInvalidValue
	}
	return nil
}

func (field *Field) GetValue() (Value, error) {
	// Ensure that at most one value exists, and that it matches the type
	if field.Type != TYPE_TEXT && field.Text != nil ||
		field.Type != TYPE_NUMBER && field.Number != nil ||
		field.Type != TYPE_CHECKBOXES && field.Checkboxes != nil ||
		field.Type != TYPE_RADIOBUTTONS && field.Radiobuttons != nil ||
		field.Type != TYPE_IMAGES && field.Images != nil {
		return nil, ErrMultipleValues
	}

	switch field.Type {
	case TYPE_TEXT:
		return field.Text, nil
	case TYPE_NUMBER:
		return field.Number, nil
	case TYPE_CHECKBOXES:
		return field.Checkboxes, nil
	case TYPE_RADIOBUTTONS:
		return field.Radiobuttons, nil
	case TYPE_IMAGES:
		return field.Images, nil
	default:
		return nil, ErrInvalidType
	}
}
