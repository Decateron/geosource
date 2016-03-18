package transactions

import (
	"testing"

	"github.com/joshheinrichs/geosource/server/types"
	"github.com/stretchr/testify/assert"
)

func TestAddUser(t *testing.T) {
	newUser := types.User{
		ID:     "u1-foo",
		Name:   "u1-bar",
		Avatar: "u1-baz",
		Email:  "u1-zap",
	}
	assert.NoError(t, AddUser(&newUser))
	assert.Error(t, AddUser(&newUser))
}

func TestGetUserByEmail(t *testing.T) {
	newUser := types.User{
		ID:     "u2-foo",
		Name:   "u2-bar",
		Avatar: "u2-baz",
		Email:  "u2-zap",
	}
	assert.NoError(t, AddUser(&newUser))

	dbUser, err := GetUserByEmail(newUser.Email)
	assert.NoError(t, err)
	assert.Equal(t, &newUser, dbUser)

	dbUser, err = GetUserByEmail("u2-asdf")
	assert.NoError(t, err)
	assert.Nil(t, dbUser)
}

func TestGetUserByID(t *testing.T) {
	newUser := types.User{
		ID:     "u3-foo",
		Name:   "u3-bar",
		Avatar: "u3-baz",
		Email:  "u3-zap",
	}
	assert.NoError(t, AddUser(&newUser))

	dbUser, err := GetUserByID(newUser.ID)
	assert.NoError(t, err)
	assert.Equal(t, &newUser, dbUser)

	dbUser, err = GetUserByEmail("u3-asdf")
	assert.NoError(t, err)
	assert.Nil(t, dbUser)
}
