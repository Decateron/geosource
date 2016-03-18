package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocationValidate(t *testing.T) {
	location := Location{
		Latitude:  -90,
		Longitude: -180,
	}
	assert.NoError(t, location.Validate())

	location = Location{
		Latitude:  90,
		Longitude: 180,
	}
	assert.NoError(t, location.Validate())

	location = Location{
		Latitude:  -91,
		Longitude: 0,
	}
	assert.Error(t, location.Validate())

	location = Location{
		Latitude:  91,
		Longitude: 0,
	}
	assert.Error(t, location.Validate())

	location = Location{
		Latitude:  0,
		Longitude: -181,
	}
	assert.Error(t, location.Validate())

	location = Location{
		Latitude:  0,
		Longitude: 181,
	}
	assert.Error(t, location.Validate())
}
