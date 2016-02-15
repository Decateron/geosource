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

func TestTextValidate(t *testing.T) {
	form := TextForm{}

	var value TextValue = "hello"
	err := form.Validate(&value)
	assert.NoError(t, err)

	value = ""
	err = form.Validate(&value)
	assert.NoError(t, err)

	err = form.Validate(nil)
	assert.Error(t, err)
}

func TestTextIsComplete(t *testing.T) {
	var value TextValue = "hello"
	var ptr *TextValue = &value
	assert.True(t, ptr.IsComplete())
	value = ""
	assert.False(t, ptr.IsComplete())
	ptr = nil
	assert.False(t, ptr.IsComplete())
}
