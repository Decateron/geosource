package fields

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

const (
	TYPE_TEXT         = "text"
	TYPE_CHECKBOXES   = "checkboxes"
	TYPE_RADIOBUTTONS = "radiobuttons"
	TYPE_IMAGES       = "images"
)

type Field struct {
	Label    string `json:"label"`
	Type     string `json:"type"`
	Required bool   `json:"required"`
	Form     Form   `json:"form"`
	Value    Value  `json:"value,omitempty"`
}

func (field *Field) IsEmpty() bool {
	return field.Value == nil
}

func (field *Field) ValidateForm() error {
	return field.Form.ValidateForm()
}

func (field *Field) ValidateValue() error {
	return field.Form.ValidateValue(field.Value)
}

type Fields []*Field

func (fields *Fields) ValidateForms() error {
	for _, field := range *fields {
		err := field.ValidateForm()
		if err != nil {
			return err
		}
	}
	return nil
}

func (fields *Fields) ValidateValues() error {
	for _, field := range *fields {
		err := field.ValidateValue()
		if err != nil {
			return err
		}
	}
	return nil
}

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

func (fields Fields) Value() (driver.Value, error) {
	blob, err := json.Marshal(fields)
	return string(blob), err
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
type Value interface {
	IsComplete() bool
}

// // Returns an error if the given field is invalid, such as in the case when
// // the form and value types are mismatched, nil otherwise.
// func (field *Field) Validate() error {
// 	if field.Value == nil {
// 		return nil
// 	}
// 	err := field.Form.Validate()
// 	if err != nil {
// 		return err
// 	}
// 	err = field.Form.ValidateValue(field.Value)
// 	if err != nil {
// 		return err
// 	}
// 	log.Println("validated")
// 	return nil
// }

// Attempts to unmarshal the given JSON into a field. Returns an error if
// unsuccessful.
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

// Attempts to unmarshal the given JSON into a form of the type specified by the
// given string. Returns an error if an invalid type is given, or the JSON
// cannot be unmarshaled into the corresponding form type.
func UnmarshalForm(fieldType string, blob []byte) (Form, error) {
	switch fieldType {
	case TYPE_TEXT:
		return &TextForm{}, nil
	case TYPE_CHECKBOXES:
		if len(blob) <= 0 {
			return nil, errors.New("No form provided for checkboxes field.")
		}
		var checkboxesForm CheckboxesForm
		err := json.Unmarshal(blob, &checkboxesForm)
		if err != nil {
			return nil, err
		}
		return &checkboxesForm, nil
	case TYPE_RADIOBUTTONS:
		if len(blob) <= 0 {
			return nil, errors.New("No form provided for radiobuttons field.")
		}
		var radiobuttonsForm RadiobuttonsForm
		err := json.Unmarshal(blob, &radiobuttonsForm)
		if err != nil {
			return nil, err
		}
		return &radiobuttonsForm, nil
	case TYPE_IMAGES:
		return &ImagesForm{}, nil
	default:
		return nil, errors.New("Invalid type.")
	}
}
