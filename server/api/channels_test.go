package api

import (
	"testing"

	"github.com/ant0ine/go-json-rest/rest/test"
	"github.com/stretchr/testify/assert"
)

func TestGetChannels(t *testing.T) {
	handler, err := MakeHandler()
	assert.NoError(t, err)
	recorded := test.RunRequest(t, handler, test.MakeSimpleRequest("GET", "http://1.2.3.4/channels", nil))
	recorded.CodeIs(200)
	recorded.ContentTypeIsJson()
}
