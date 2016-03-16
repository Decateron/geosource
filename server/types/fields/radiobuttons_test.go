package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRadiobuttonsUnmarshalValue(t *testing.T) {
	form := RadiobuttonsForm{}
	data := `"foo"`
	_, err := form.UnmarshalValue([]byte(data))
	assert.NoError(t, err)

	data = `[]`
	_, err = form.UnmarshalValue([]byte(data))
	assert.Error(t, err)

	data = ``
	_, err = form.UnmarshalValue([]byte(data))
	assert.NoError(t, err)

	data = `{}`
	_, err = form.UnmarshalValue([]byte(data))
	assert.Error(t, err)

	data = `{"value": "foo"}`
	_, err = form.UnmarshalValue([]byte(data))
	assert.Error(t, err)
}

func TestRadiobuttonsValidateForm(t *testing.T) {
	form := &RadiobuttonsForm{}
	assert.Error(t, form.ValidateForm())

	form = &RadiobuttonsForm{"foo", "bar", "foo"}
	assert.Error(t, form.ValidateForm())

	form = nil
	assert.Error(t, form.ValidateForm())

	form = &RadiobuttonsForm{"foo", "bar"}
	assert.NoError(t, form.ValidateForm())

	form = &RadiobuttonsForm{"foo"}
	assert.NoError(t, form.ValidateForm())
}

func TestRadiobuttonsValidateValue(t *testing.T) {
	form := RadiobuttonsForm{"foo", "bar"}
	var value RadiobuttonsValue
	ptr := &value

	value = "foo"
	assert.NoError(t, form.ValidateValue(ptr))

	value = "bar"
	assert.NoError(t, form.ValidateValue(ptr))

	value = "baz"
	assert.Error(t, form.ValidateValue(ptr))

	ptr = nil
	assert.NoError(t, form.ValidateValue(ptr))

	assert.Error(t, form.ValidateValue(nil))
}

func TestRadiobuttonsIsComplete(t *testing.T) {
	var value RadiobuttonsValue = ""
	ptr := &value
	assert.True(t, ptr.IsComplete())
	ptr = nil
	assert.False(t, ptr.IsComplete())
}
