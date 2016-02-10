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
	Value    interface{} `json:"value"`
}

type Form interface {
	Validate(Value) error
}
type Value interface {
	IsComplete() bool
}

func (field *Field) Validate() error {
	form, ok := field.Form.(Form)
	if !ok {
		return errors.New("Invalid form type.")
	}
	value, ok := field.Value.(Value)
	if !ok {
		return errors.New("Invalid value type.")
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

	switch field.Type {
	case TYPE_CHECKBOXES:
		var checkboxesForm CheckboxesForm
		err = json.Unmarshal(jsonForm, &checkboxesForm)
		if err != nil {
			return nil, err
		}
		var checkboxesValue CheckboxesValue
		err = json.Unmarshal(jsonValue, &checkboxesValue)
		if err != nil {
			return nil, err
		}
		field.Form = checkboxesForm
		field.Value = checkboxesValue
	case TYPE_TEXT:
		var textForm TextForm
		err = json.Unmarshal(jsonForm, &textForm)
		if err != nil {
			return nil, err
		}
		var textValue TextValue
		err = json.Unmarshal(jsonValue, &textValue)
		if err != nil {
			return nil, err
		}
		field.Form = textForm
		field.Value = textValue
	default:
		return nil, errors.New("Invalid type.")
	}

	return &field, nil
}
