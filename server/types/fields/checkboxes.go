package fields

import (
	"encoding/json"
	"errors"
)

type CheckboxesForm []string

func (checkboxesForm *CheckboxesForm) Validate() error {
	if checkboxesForm == nil {
		return errors.New("Missing form.")
	}
	if len(*checkboxesForm) <= 0 {
		return errors.New("At least one checkbox is required.")
	}
	return nil
}

func (checkboxesForm *CheckboxesForm) ValidateValue(value Value) error {
	checkboxesValue, ok := value.(*CheckboxesValue)
	if !ok {
		return errors.New("Type mismatch.")
	}
	if checkboxesValue == nil {
		return nil
	}
	if len(*checkboxesForm) != len(*checkboxesValue) {
		return errors.New("Length mismatch.")
	}
	return nil
}

func (checkboxesForm *CheckboxesForm) UnmarshalValue(blob []byte) (Value, error) {
	if len(blob) <= 0 {
		return nil, nil
	}
	var value CheckboxesValue
	err := json.Unmarshal(blob, &value)
	if err != nil {
		return nil, err
	}
	return &value, nil
}

type CheckboxesValue []bool

func (checkboxesValue *CheckboxesValue) IsComplete() bool {
	return checkboxesValue != nil
}
