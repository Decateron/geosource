package main

import (
	"encoding/json"
	"errors"
)

type CheckboxesForm []string

func (checkboxesForm *CheckboxesForm) Validate(value Value) error {
	checkboxesValue, ok := value.(*CheckboxesValue)
	if !ok {
		return errors.New("Type mismatch.")
	}
	if len(*checkboxesForm) != len(*checkboxesValue) {
		return errors.New("Length mismatch.")
	}
	return nil
}

func (checkboxesForm *CheckboxesForm) UnmarshalValue(blob []byte) (Value, error) {
	if len(blob) > 0 {
		var value CheckboxesValue
		err := json.Unmarshal(blob, &value)
		if err != nil {
			return nil, err
		}
		return &value, nil
	}
	return nil, nil
}

type CheckboxesValue []bool

func (checkboxesValue *CheckboxesValue) IsComplete() bool {
	return true
}
