package field

import (
	"errors"
)

const (
	FIELDTYPE_TEXT         string = "Text"
	FIELDTYPE_CHECKBOXES   string = "CheckBoxes"
	FIELDTYPE_CHECKBOX     string = "CheckBox"
	FIELDTYPE_RADIOBUTTONS string = "RadioButtons"
	FIELDTYPE_RADIOBUTTON  string = "RadioButton"
	FIELDTYPE_IMAGES       string = "Images"
)

var ErrInvalidType = errors.New("Invalid field type")
var ErrInvalidValue = errors.New("Invalid value type")

type Field struct {
	Label *string     `json:"label"`
	Type  *string     `json:"type"`
	Value interface{} `json:"value"`
}

func ValidField(field *Field) (bool, error) {
	switch *field.Type {
	case FIELDTYPE_CHECKBOXES:
		return ValidCheckBoxes(field)
	case FIELDTYPE_RADIOBUTTONS:
		return ValidRadioButton(field)
	case FIELDTYPE_TEXT:
		return ValidText(field)
	case FIELDTYPE_IMAGES:
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
	if *field.Type != FIELDTYPE_IMAGES {
		return false, ErrInvalidType
	}
	_, ok := field.Value.(*string)
	if !ok {
		return false, ErrInvalidValue
	}
	return true, nil
}
