package fields

import (
	"encoding/json"
	"errors"
)

const (
	TYPE_TEXT       = "text"
	TYPE_CHECKBOXES = "checkboxes"
)

type Field struct {
	Label    string `json:"label"`
	Type     string `json:"type"`
	Required bool   `json:"required"`
	Form     Form   `json:"form"`
	Value    Value  `json:"value,omitempty"`
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
	if field.Value == nil {
		return nil
	}
	err := field.Form.Validate(field.Value)
	if err != nil {
		return err
	}
	return nil
}

func UnmarshalField(blob []byte) (*Field, error) {
	unmarshalField := struct {
		Field
		JsonForm  json.RawMessage `json:"form"`
		JsonValue json.RawMessage `json:"value"`
	}{}
	err := json.Unmarshal(blob, &unmarshalField)
	if err != nil {
		return nil, err
	}
	form, err := UnmarshalForm(unmarshalField.Type, unmarshalField.JsonForm)
	if err != nil {
		return nil, err
	}
	value, err := form.UnmarshalValue(unmarshalField.JsonValue)
	if err != nil {
		return nil, err
	}
	unmarshalField.Field.Form = form
	unmarshalField.Field.Value = value
	return &unmarshalField.Field, nil
}

func UnmarshalForm(fieldType string, blob []byte) (Form, error) {
	switch fieldType {
	case TYPE_CHECKBOXES:
		if len(blob) > 0 {
			var checkboxesForm CheckboxesForm
			err := json.Unmarshal(blob, &checkboxesForm)
			if err != nil {
				return nil, err
			}
			return &checkboxesForm, nil
		} else {
			return nil, errors.New("No form provided for checkboxes field.")
		}
	case TYPE_TEXT:
		return &TextForm{}, nil
	default:
		return nil, errors.New("Invalid type.")
	}
}
