package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalField(t *testing.T) {
	textJSON := `{"label": "foo", "type": "text", "required": true, "value": "bar"}`
	text, err := UnmarshalField([]byte(textJSON))
	assert.NoError(t, err)
	assert.Equal(t, "foo", text.Label)
	assert.Equal(t, "text", text.Type)
	assert.Equal(t, true, text.Required)
	assert.Equal(t, TextForm{}, *text.Form.(*TextForm))
	assert.Equal(t, "bar", string(*text.Value.(*TextValue)))

	checkboxesJSON := `{"label": "foo", "type": "checkboxes", "required": true, "form": ["foo", "bar"], "value": [true, false]}`
	checkboxes, err := UnmarshalField([]byte(checkboxesJSON))
	assert.NoError(t, err)
	assert.Equal(t, "foo", checkboxes.Label)
	assert.Equal(t, "checkboxes", checkboxes.Type)
	assert.Equal(t, true, checkboxes.Required)
	assert.Equal(t, CheckboxesForm{"foo", "bar"}, *checkboxes.Form.(*CheckboxesForm))
	assert.Equal(t, CheckboxesValue{true, false}, *checkboxes.Value.(*CheckboxesValue))

	radiobuttonsJSON := `{"label": "foo", "type": "radiobuttons", "required": true, "form": ["foo", "bar"], "value": "foo"}`
	radiobuttons, err := UnmarshalField([]byte(radiobuttonsJSON))
	assert.NoError(t, err)
	assert.Equal(t, "foo", radiobuttons.Label)
	assert.Equal(t, "radiobuttons", radiobuttons.Type)
	assert.Equal(t, true, radiobuttons.Required)
	assert.Equal(t, RadiobuttonsForm{"foo", "bar"}, *radiobuttons.Form.(*RadiobuttonsForm))
	assert.Equal(t, "foo", string(*radiobuttons.Value.(*RadiobuttonsValue)))

	imagesJSON := `{"label": "foo", "type": "images", "required": false, "value": ["foo", "bar"]}`
	images, err := UnmarshalField([]byte(imagesJSON))
	assert.NoError(t, err)
	assert.Equal(t, "foo", images.Label)
	assert.Equal(t, "images", images.Type)
	assert.Equal(t, false, images.Required)
	assert.Equal(t, ImagesForm{}, *images.Form.(*ImagesForm))
	assert.Equal(t, ImagesValue{"foo", "bar"}, *images.Value.(*ImagesValue))

	errJSON := `{"type": "error"}`
	_, err = UnmarshalField([]byte(errJSON))
	assert.Error(t, err)
}
