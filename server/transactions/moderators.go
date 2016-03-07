package transactions

func AddModerator(requesterUid, uid, channelname string) error {
	return nil
}

func GetModerators(requesterUid, channelname string) ([]string, error) {
	return nil, nil
}

func IsModerator(requesterUid, uid, channelname string) (bool, error) {
	return false, nil
}

func RemoveModerator(requesterUid, uid, channelname string) error {
	return nil
}
