package field

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTextValidate(t *testing.T) {
	var encoded string
	var field Field
	var err error

	// Should succeed
	field = Field{}
	encoded = `
	{
		"type": "text",
		"label": "foo",
		"required": true,
		"text": "bar"
	}
	`
	err = json.Unmarshal([]byte(encoded), &field)
	assert.NoError(t, err)
	assert.NoError(t, field.Validate())

	// Empty text when required
	field = Field{}
	encoded = `
	{
		"type": "text",
		"label": "foo",
		"required": true,
		"text": ""
	}
	`
	err = json.Unmarshal([]byte(encoded), &field)
	assert.NoError(t, err)
	assert.Error(t, field.Validate())

	// Missing text when required
	field = Field{}
	encoded = `
	{
		"type": "text",
		"label": "foo",
		"required": true
	}
	`
	err = json.Unmarshal([]byte(encoded), &field)
	assert.NoError(t, err)
	assert.Error(t, field.Validate())

	// Additional field
	field = Field{}
	encoded = `
	{
		"type": "text",
		"label": "foo",
		"required": true,
		"text": "bar",
		"images": ["baz"]
	}
	`
	err = json.Unmarshal([]byte(encoded), &field)
	assert.NoError(t, err)
	assert.Error(t, field.Validate())
}
