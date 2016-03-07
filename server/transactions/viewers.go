package transactions

import "errors"

func AddViewer(requesterUid, uid, channelname string) error {
	return errors.New("function has not yet been implemented.")
}

func GetViewers(requesterUid, channelname string) ([]string, error) {
	return nil, errors.New("function has not yet been implemented.")
}

func IsViewer(requesterUid, uid, channelname string) (bool, error) {
	return false, errors.New("function has not yet been implemented.")
}

func RemoveViewer(requesterUid, uid, channelname string) error {
	return errors.New("function has not yet been implemented.")
}
