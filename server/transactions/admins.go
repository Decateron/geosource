package transactions

import (
	"errors"
)

func AddAdmin(requesterID, userID string) error {
	return errors.New("function has not yet been implemented")
}

func IsAdmin(userID string) (bool, error) {
	return false, errors.New("function has not yet been implemented")
}

func RemoveAdmin(requesterID, userID string) error {
	return errors.New("function has not yet been implemented")
}
