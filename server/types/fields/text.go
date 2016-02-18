package fields

import (
	"encoding/json"
	"errors"
	"strings"
)

type TextForm struct{}

func (textForm *TextForm) Validate() error {
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

type TextValue string

func (textValue *TextValue) IsComplete() bool {
	return textValue != nil && len(strings.TrimSpace(string(*textValue))) > 0
}
