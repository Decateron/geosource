package main

import (
	"errors"
)

type CheckboxesForm []string

func (checkboxesForm CheckboxesForm) Validate(value Value) error {
	checkboxesValue, ok := value.(CheckboxesValue)
	if !ok {
		return errors.New("Type mismatch.")
	}
	if len(checkboxesForm) != len(checkboxesValue) {
		return errors.New("Length mismatch.")
	}
	return nil
}

type CheckboxesValue []bool

func (checkboxesValue CheckboxesValue) IsComplete() bool {
	return true
}
