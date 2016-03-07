package transactions

import "errors"

func AddModerator(requesterUid, uid, channelname string) error {
	return errors.New("function has not yet been implemented.")
}

func GetModerators(requesterUid, channelname string) ([]string, error) {
	return nil, errors.New("function has not yet been implemented.")
}

func IsModerator(requesterUid, uid, channelname string) (bool, error) {
	return false, errors.New("function has not yet been implemented.")
}

func RemoveModerator(requesterUid, uid, channelname string) error {
	return errors.New("function has not yet been implemented.")
}
