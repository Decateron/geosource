package fields

import (
	"encoding/json"
	"errors"
	"strings"
)

type TextForm struct{}

func (textForm *TextForm) Validate(value Value) error {
	_, ok := value.(*TextValue)
	if !ok {
		return errors.New("Type mismatch.")
	}
	return nil
}

func (textForm *TextForm) UnmarshalValue(blob []byte) (Value, error) {
	if len(blob) > 0 {
		var value TextValue
		err := json.Unmarshal(blob, &value)
		if err != nil {
			return nil, err
		}
		return &value, nil
	}
	return nil, nil
}

type TextValue string

func (textValue *TextValue) IsComplete() bool {
	return len(strings.TrimSpace(string(*textValue))) > 0
}
