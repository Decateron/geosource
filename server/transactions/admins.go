package transactions

import (
	"errors"
)

func AddAdmin(requester, userid string) error {
	return errors.New("function has not yet been implemented.")
}

func IsAdmin(userid string) (bool, error) {
	return false, errors.New("function has not yet been implemented.")
}

func RemoveAdmin(requester, userid string) error {
	return errors.New("function has not yet been implemented.")
}
