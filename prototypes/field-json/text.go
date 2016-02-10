package main

import (
	"errors"
	"strings"
)

type TextForm struct{}

func (textForm TextForm) Validate(value Value) error {
	_, ok := value.(TextValue)
	if !ok {
		return errors.New("Type mismatch.")
	}
	return nil
}

type TextValue string

func (textValue TextValue) IsComplete() bool {
	return len(strings.TrimSpace(string(textValue))) > 0
}
