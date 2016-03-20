package fields

import (
	"encoding/json"
	"errors"
	"strings"
)

// No form is needed for images as there are no limitations that can be set by
// the user.
type TextForm struct{}

func (textForm *TextForm) ValidateForm() error {
	return nil
}

func (textForm *TextForm) ValidateValue(value Value) error {
	_, ok := value.(*TextValue)
	if !ok {
		return errors.New("Type mismatch.")
	}
	return nil
}

func (textForm *TextForm) UnmarshalValue(blob []byte) (Value, error) {
	if len(blob) <= 0 {
		return nil, nil
	}
	var value TextValue
	err := json.Unmarshal(blob, &value)
	if err != nil {
		return nil, err
	}
	return &value, nil
}

// A text value is simply a string, representing the text input by the user.
type TextValue string

func (textValue *TextValue) IsEmpty() bool {
	return !textValue.IsComplete()
}

func (textValue *TextValue) IsComplete() bool {
	return textValue != nil && len(strings.TrimSpace(string(*textValue))) > 0
}
