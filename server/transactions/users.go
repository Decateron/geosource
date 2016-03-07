package transactions

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/joshheinrichs/geosource/server/types"
)

// Adds the given User to the database, or returns an error if the insertion
// was unsuccessful.
func AddUser(user *types.User) error {
	err := db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUserByEmail(email string) (*types.User, error) {
	var user types.User
	err := db.Where("u_email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.RecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func GetUserById(id string) (*types.User, error) {
	var user types.User
	err := db.Where("u_userid = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUsername(email string) (*string, error) {
	return nil, errors.New("function has not yet been implemented.")
}

// Returns a user with the given username or nil if they do not exist. An error
// is returned if the database was accessed unsuccessfully.
func GetUser(requesterUid, uid string) (*types.User, error) {
	return nil, errors.New("function has not yet been implemented.")
}

func RemoveUser(requesterUid, uid string) error {
	return errors.New("function has not yet been implemented.")
}
