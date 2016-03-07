package transactions

import (
	"errors"
)

func AddBan(requesterUid, uid, channelname string) error {
	return errors.New("function has not yet been implemented.")
}

func GetBans(requesterUid, channelname string) ([]string, error) {
	return nil, errors.New("function has not yet been implemented.")
}

func IsBanned(requesterUid, uid, channelname string) (bool, error) {
	return false, errors.New("function has not yet been implemented.")
}

func RemoveBan(requesterUid, uid, channelname string) error {
	return errors.New("function has not yet been implemented.")
}
