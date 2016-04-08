package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTextUnmarshalValue(t *testing.T) {
	form := TextForm{}
	data := `"hello"`
	_, err := form.UnmarshalValue([]byte(data))
	assert.NoError(t, err)

	data = ``
	_, err = form.UnmarshalValue([]byte(data))
	assert.NoError(t, err)

	data = `{}`
	_, err = form.UnmarshalValue([]byte(data))
	assert.Error(t, err)

	data = `[]`
	_, err = form.UnmarshalValue([]byte(data))
	assert.Error(t, err)
}

func TestTextValidateForm(t *testing.T) {
	form := &TextForm{}
	assert.NoError(t, form.ValidateForm())

	form = nil
	assert.NoError(t, form.ValidateForm())
}

func TestTextValidateValue(t *testing.T) {
	form := TextForm{}
	var value TextValue
	ptr := &value

	value = "hello"
	assert.NoError(t, form.ValidateValue(ptr))

	value = ""
	assert.NoError(t, form.ValidateValue(ptr))

	ptr = nil
	assert.NoError(t, form.ValidateValue(ptr))

	assert.Error(t, form.ValidateValue(nil))
}

func TestTextIsComplete(t *testing.T) {
	var value TextValue = "hello"
	var ptr = &value
	assert.True(t, ptr.IsComplete())
	value = ""
	assert.False(t, ptr.IsComplete())
	ptr = nil
	assert.False(t, ptr.IsComplete())
}
