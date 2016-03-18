package transactions

import (
	"log"
	"testing"

	"github.com/joshheinrichs/geosource/server/types"
	"github.com/stretchr/testify/assert"
)

func TestAddUser(t *testing.T) {
	newUser := types.User{
		ID:     "foo",
		Name:   "bar",
		Avatar: "baz",
		Email:  "zap",
	}
	assert.NoError(t, AddUser(&newUser))
	assert.Error(t, AddUser(&newUser))
}

func TestGetUserByEmail(t *testing.T) {
	newUser := types.User{
		ID:     "foo",
		Name:   "bar",
		Avatar: "baz",
		Email:  "zap",
	}
	AddUser(&newUser) // Add user if it doesn't exist
	dbUser, err := GetUserByEmail(newUser.Email)
	assert.NoError(t, err)
	assert.Equal(t, &newUser, dbUser)
	log.Println("test")
}
