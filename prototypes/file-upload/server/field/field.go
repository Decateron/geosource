package field

import (
	"errors"
)

const (
	FIELDTYPE_TEXT         string = "text"
	FIELDTYPE_CHECKBOXES   string = "checkboxes" 
	FIELDTYPE_CHECKBOX     string = "checkbox"
	FIELDTYPE_RADIOBUTTONS string = "radiobuttons"
	FIELDTYPE_RADIOBUTTON  string = "radiobutton"
	FIELDTYPE_IMAGE        string = "image"
)

var ErrInvalidType = errors.New("Invalid field type")
var ErrInvalidValue = errors.New("Invalid value type")

type Field struct {
	Label     *string     `json:"label"`
	Type      *string     `json:"type"`
	Value     interface{} `json:"value"`
}

func ValidField(field *Field) (bool, error) {
	switch *field.Type {
	case FIELDTYPE_CHECKBOXES:
		return ValidCheckBoxes(field)
	case FIELDTYPE_RADIOBUTTONS:
		return ValidRadioButton(field)
	case FIELDTYPE_TEXT:
		return ValidText(field)
	case FIELDTYPE_IMAGE:
		return ValidImage(field)
	default:
		return false, ErrInvalidType
	}
}

func ValidText(field *Field) (bool, error) {
	if *field.Type != FIELDTYPE_TEXT {
		return false, ErrInvalidType
	} 
	_, ok := field.Value.(*string)
	if !ok {
		return false, ErrInvalidValue
	}
	return true, nil
}

func ValidCheckBoxes(field *Field) (bool, error) {
	if *field.Type != FIELDTYPE_CHECKBOXES {
		return false, ErrInvalidType
	}
	_, ok := field.Value.([]Field)
	if !ok {
		return false, ErrInvalidValue
	}
	return true, nil
}

func ValidCheckBox(field *Field) (bool, error) {
	if *field.Type != FIELDTYPE_CHECKBOXES {
		return false, ErrInvalidType
	}
	_, ok := field.Value.(*bool)
	if !ok {
		return false, ErrInvalidValue
	}
	return true, nil
}

func ValidRadioButtons(field *Field) (bool, error) {
	if *field.Type != FIELDTYPE_RADIOBUTTONS {
		return false, ErrInvalidType
	}
	_, ok := field.Value.([]Field)
	if !ok {
		return false, ErrInvalidValue
	}
	return true, nil
}

func ValidRadioButton(field *Field) (bool, error) {
	if *field.Type != FIELDTYPE_RADIOBUTTON {
		return false, ErrInvalidType
	}
	_, ok := field.Value.(*bool)
	if !ok {
		return false, ErrInvalidValue
	}
	return true, nil
}

func ValidImage(field *Field) (bool, error) {
	if *field.Type != FIELDTYPE_IMAGE {
		return false, ErrInvalidType
	}
	_, ok := field.Value.(*string)
	if !ok {
		return false, ErrInvalidValue
	}
	return true, nil
}
