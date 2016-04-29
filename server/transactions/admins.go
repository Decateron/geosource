package transactions

import "github.com/joshheinrichs/httperr"

func AddAdmin(requesterID, userID string) httperr.Error {
	return ErrNotImplemented
}

func IsAdmin(userID string) (bool, httperr.Error) {
	return false, ErrNotImplemented
}

func RemoveAdmin(requesterID, userID string) httperr.Error {
	return ErrNotImplemented
}
