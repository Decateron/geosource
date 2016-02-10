package main

import (
	"encoding/json"
	"errors"
)

const (
	TYPE_TEXT       = "text"
	TYPE_CHECKBOXES = "checkboxes"
)

type Field struct {
	Label    string      `json:"label"`
	Type     string      `json:"type"`
	Required bool        `json:"required"`
	Form     interface{} `json:"form"`
	Value    interface{} `json:"value,omitempty"`
}

type Form interface {
	Validate(Value) error
	UnmarshalValue([]byte) (Value, error)
}
type Value interface {
	IsComplete() bool
}

func (field *Field) IsEmpty() bool {
	return field.Value == nil
}

func (field *Field) Validate() error {
	form, ok := field.Form.(Form)
	if !ok {
		return errors.New("Invalid form.")
	}
	if field.Value == nil {
		return nil
	}
	value, ok := field.Value.(Value)
	if !ok {
		return errors.New("Invalid value.")
	}
	err := form.Validate(value)
	if err != nil {
		return err
	}
	return nil
}

func UnmarshalField(blob []byte) (*Field, error) {
	var jsonForm json.RawMessage
	var jsonValue json.RawMessage
	field := Field{
		Form:  &jsonForm,
		Value: &jsonValue,
	}
	err := json.Unmarshal(blob, &field)
	if err != nil {
		return nil, err
	}

	var form Form
	switch field.Type {
	case TYPE_CHECKBOXES:
		if len(jsonForm) > 0 {
			var checkboxesForm CheckboxesForm
			err = json.Unmarshal(jsonForm, &checkboxesForm)
			if err != nil {
				return nil, err
			}
			form = &checkboxesForm
		} else {
			form = &CheckboxesForm{}
		}
		// if len(jsonValue) > 0 {
		// 	var checkboxesValue CheckboxesValue
		// 	err = json.Unmarshal(jsonValue, &checkboxesValue)
		// 	if err != nil {
		// 		return nil, err
		// 	}
		// 	field.Value = checkboxesValue
		// } else {
		// 	field.Value = nil
		// }
	case TYPE_TEXT:
		form = &TextForm{}
		// if len(jsonValue) > 0 {
		// 	var textValue TextValue
		// 	err = json.Unmarshal(jsonValue, &textValue)
		// 	if err != nil {
		// 		return nil, err
		// 	}
		// 	field.Value = textValue
		// } else {
		// 	field.Value = nil
		// }
	default:
		return nil, errors.New("Invalid type.")
	}

	value, err := form.UnmarshalValue(jsonValue)
	if err != nil {
		return nil, err
	}
	field.Form = form
	field.Value = value

	return &field, nil
}
