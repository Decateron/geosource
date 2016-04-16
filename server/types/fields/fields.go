package fields

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/joshheinrichs/geosource/server/config"
)

const (
	TypeText         = "text"
	TypeCheckboxes   = "checkboxes"
	TypeRadiobuttons = "radiobuttons"
	TypeImages       = "images"
)

var fieldsConfig *config.Config

func Init(config *config.Config) {
	fieldsConfig = config
}

type Field struct {
	Label    string `json:"label"`
	Type     string `json:"type"`
	Required bool   `json:"required"`
	Form     Form   `json:"form"`
	Value    Value  `json:"value,omitempty"`
}

// UnmarshalField attempts to unmarshal the given JSON into a field. Returns an
// error if unsuccessful.
func UnmarshalField(blob []byte) (*Field, error) {
	unmarshalField := struct {
		Field
		JSONForm  json.RawMessage `json:"form"`
		JSONValue json.RawMessage `json:"value"`
	}{}
	err := json.Unmarshal(blob, &unmarshalField)
	if err != nil {
		return nil, err
	}
	form, err := UnmarshalForm(unmarshalField.Type, unmarshalField.JSONForm)
	if err != nil {
		return nil, err
	}
	value, err := form.UnmarshalValue(unmarshalField.JSONValue)
	if err != nil {
		return nil, err
	}
	unmarshalField.Field.Form = form
	unmarshalField.Field.Value = value
	return &unmarshalField.Field, nil
}

func (field *Field) IsEmpty() bool {
	return field.Value == nil || field.Value.IsEmpty()
}

func (field *Field) ValidateForm() error {
	return field.Form.ValidateForm()
}

func (field *Field) ValidateValue() error {
	if field.Required && field.IsEmpty() {
		return errors.New("Required field is empty")
	} else if field.IsEmpty() {
		return nil
	}
	return field.Form.ValidateValue(field.Value)
}

type Form interface {
	// Returns an error if the form is invalid, nil otherwise.
	ValidateForm() error
	// Returns an error if the given value does not match the form, nil
	// otherwise.
	ValidateValue(Value) error
	// Attempts to unmarshal the given JSON into this form's corresponding
	// value type. Returns an error if unsuccessful.
	UnmarshalValue([]byte) (Value, error)
}

// UnmarshalForm attempts to unmarshal the given JSON into a form of the type
// specified by the given string. Returns an error if an invalid type is given,
// or the JSON cannot be unmarshaled into the corresponding form type.
func UnmarshalForm(fieldType string, blob []byte) (Form, error) {
	switch fieldType {
	case TypeText:
		return &TextForm{}, nil
	case TypeCheckboxes:
		if len(blob) <= 0 {
			return nil, errors.New("No form provided for checkboxes field.")
		}
		var checkboxesForm CheckboxesForm
		err := json.Unmarshal(blob, &checkboxesForm)
		if err != nil {
			return nil, err
		}
		return &checkboxesForm, nil
	case TypeRadiobuttons:
		if len(blob) <= 0 {
			return nil, errors.New("No form provided for radiobuttons field.")
		}
		var radiobuttonsForm RadiobuttonsForm
		err := json.Unmarshal(blob, &radiobuttonsForm)
		if err != nil {
			return nil, err
		}
		return &radiobuttonsForm, nil
	case TypeImages:
		return &ImagesForm{}, nil
	default:
		return nil, errors.New("Invalid type.")
	}
}

type Value interface {
	// IsEmpty returns true if the value is empty, false otherwise.
	IsEmpty() bool
	// IsComplete returns true if the value is complete, false otherwise. A
	// value could be both non-empty and incomplete making both of these methods
	// necessary.
	IsComplete() bool
}

// Fields is an array of Field. This allows concise validation and easy
// insertion and retrieval from the database given its current structure.
type Fields []*Field

// ValidateForms returns an error if any of the fields' forms are invalid, nil
// otherwise.
func (fields *Fields) ValidateForms() error {
	for _, field := range *fields {
		err := field.ValidateForm()
		if err != nil {
			return err
		}
	}
	return nil
}

// ValidateValues returns an error if any of the fields' values are invalid,
// nil otherwise.
func (fields *Fields) ValidateValues() error {
	for _, field := range *fields {
		err := field.ValidateValue()
		if err != nil {
			return err
		}
	}
	return nil
}

// Scan converts a set of fields from their database representation (JSON) into
// the fields structure. This allows for automatic unparsing in transactions.
func (fields *Fields) Scan(value interface{}) error {
	blob, ok := value.([]byte)
	if !ok {
		return errors.New("Could not convert interface to byte array.")
	}
	var jsonFields []json.RawMessage
	err := json.Unmarshal(blob, &jsonFields)
	if err != nil {
		return err
	}
	*fields = make([]*Field, len(jsonFields))
	for i, jsonField := range jsonFields {
		field, err := UnmarshalField(jsonField)
		if err != nil {
			return err
		}
		(*fields)[i] = field
	}
	return nil
}

// Value converts a set of fields into their corresponding database
// representation (JSON). This allows for automatic parsing in transactions.
func (fields Fields) Value() (driver.Value, error) {
	blob, err := json.Marshal(fields)
	return string(blob), err
}
