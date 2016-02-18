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

func TestRadiobuttonsValidate(t *testing.T) {
	form := &RadiobuttonsForm{}
	assert.Error(t, form.Validate())

	form = &RadiobuttonsForm{"foo", "bar", "foo"}
	assert.Error(t, form.Validate())

	form = nil
	assert.Error(t, form.Validate())

	form = &RadiobuttonsForm{"foo", "bar"}
	assert.NoError(t, form.Validate())

	form = &RadiobuttonsForm{"foo"}
	assert.NoError(t, form.Validate())
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

// func TestRadiobuttonsIsComplete(t *testing.T) {
// 	value := &RadiobuttonsValue{}
// 	assert.True(t, value.IsComplete())
// 	value = &RadiobuttonsValue{true, false}
// 	assert.True(t, value.IsComplete())
// 	value = nil
// 	assert.False(t, value.IsComplete())
// }
