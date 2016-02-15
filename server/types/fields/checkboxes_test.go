package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckboxesUnmarshalValue(t *testing.T) {
	form := CheckboxesForm{}
	data := `[true, false]`
	_, err := form.UnmarshalValue([]byte(data))
	assert.NoError(t, err)

	data = `[]`
	_, err = form.UnmarshalValue([]byte(data))
	assert.NoError(t, err)

	data = ``
	_, err = form.UnmarshalValue([]byte(data))
	assert.NoError(t, err)

	data = `{}`
	_, err = form.UnmarshalValue([]byte(data))
	assert.Error(t, err)

	data = `{"array": [true, false]}`
	_, err = form.UnmarshalValue([]byte(data))
	assert.Error(t, err)
}

func TestCheckboxesValidate(t *testing.T) {
	form := CheckboxesForm{"foo", "bar"}

	value := &CheckboxesValue{true, false}
	err := form.Validate(value)
	assert.NoError(t, err)

	value = &CheckboxesValue{true, false, false}
	err = form.Validate(value)
	assert.Error(t, err)

	value = &CheckboxesValue{}
	err = form.Validate(value)
	assert.Error(t, err)

	err = form.Validate(nil)
	assert.Error(t, err)
}
