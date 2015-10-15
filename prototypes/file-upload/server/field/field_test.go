package field

import (
	"testing"
)

func strPtr(s string) *string {
	return &s
}

func TestValidText(t *testing.T) {
	text := &Field {
		Label: strPtr("Foo"),
		Type: strPtr("text"),
		Value: strPtr("Bar"),
	}
	valid, err := ValidText(text)
	if !valid {
		t.Error(err)
	}
}