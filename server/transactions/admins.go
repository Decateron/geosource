package transactions

import (
	"errors"
)

func AddAdmin(requesterUid, uid string) error {
	return errors.New("function has not yet been implemented.")
}

func IsAdmin(requesterUid, uid string) (bool, error) {
	return false, errors.New("function has not yet been implemented.")
}

func RemoveAdmin(requesterUid, uid string) error {
	return errors.New("function has not yet been implemented.")
}
