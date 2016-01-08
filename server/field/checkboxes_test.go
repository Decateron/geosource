package field

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckboxesValidate(t *testing.T) {
	var encoded string
	var field Field
	var err error

	// Should succeed
	field = Field{}
	encoded = `
	{
		"type": "checkboxes",
		"label": "foo",
		"required": false,
		"checkboxes": [
			{"label": "foo", "value": true},
			{"label": "bar", "value": false},
			{"label": "baz", "value": null}
		]
	}
	`
	err = json.Unmarshal([]byte(encoded), &field)
	assert.NoError(t, err)
	assert.NoError(t, field.Validate())

	// Should succeed
	field = Field{}
	encoded = `
	{
		"type": "checkboxes",
		"label": "foo",
		"required": true,
		"checkboxes": [
			{"label": "foo", "value": true},
			{"label": "bar", "value": false}
		]
	}
	`
	err = json.Unmarshal([]byte(encoded), &field)
	assert.NoError(t, err)
	assert.NoError(t, field.Validate())

	// Should fail
	field = Field{}
	encoded = `
	{
		"type": "checkboxes",
		"label": "foo",
		"required": true,
		"checkboxes": [
			{"label": "foo", "value": true},
			{"label": "bar", "value": null}
		]
	}
	`
	err = json.Unmarshal([]byte(encoded), &field)
	assert.NoError(t, err)
	assert.Error(t, field.Validate())
}
