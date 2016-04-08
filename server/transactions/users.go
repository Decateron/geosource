package transactions

import (
	"github.com/jinzhu/gorm"
	"github.com/joshheinrichs/geosource/server/types"
)

// AddUser adds the given User to the database, or returns an error if the
// insertion was unsuccessful.
func AddUser(user *types.User) error {
	err := db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

// GetUserByEmail returns the user with the given email if one exists, nil
// otherwise. Returns an error if some error occurs within the database.
func GetUserByEmail(email string) (*types.User, error) {
	var user types.User
	err := db.Where("u_email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// GetUserByID returns the user with the given ID if one exists, nil otherwise.
// Returns an error if some error occurs within the database.
func GetUserByID(userID string) (*types.User, error) {
	var user types.User
	err := db.Where("u_userid = ?", userID).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
