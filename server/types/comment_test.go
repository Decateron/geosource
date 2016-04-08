package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommentValidate(t *testing.T) {
	// Testing empty string
	comment := Comment{}
	assert.Error(t, comment.Validate())

	// Testing minimum string length
	comment.Comment = "a"
	assert.NoError(t, comment.Validate())

	// Testing maximum string length
	maxString := ""
	for i := 0; i < MaxCommentLength; i++ {
		maxString += "a"
	}
	comment.Comment = maxString
	assert.NoError(t, comment.Validate())
	assert.Equal(t, maxString, comment.Comment)

	// Testing trimming
	comment.Comment = "   " + maxString + "   "
	assert.NoError(t, comment.Validate())
	assert.Equal(t, maxString, comment.Comment)

	// Testing string that is too large
	comment.Comment = maxString + "a"
	assert.Error(t, comment.Validate())
}
