package transactions

func AddViewer(requesterUid, uid, channelname string) error {
	return nil
}

func GetViewers(requesterUid, channelname string) ([]string, error) {
	return nil, nil
}

func IsViewer(requesterUid, uid, channelname string) (bool, error) {
	return false, nil
}

func RemoveViewer(requesterUid, uid, channelname string) error {
	return nil
}
