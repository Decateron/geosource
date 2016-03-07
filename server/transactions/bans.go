package transactions

func AddBan(requesterUid, uid, channelname string) error {
	return nil
}

func GetBans(requesterUid, channelname string) ([]string, error) {
	return nil, nil
}

func IsBanned(requesterUid, uid, channelname string) (bool, error) {
	return false, nil
}

func RemoveBan(requesterUid, uid, channelname string) error {
	return nil
}
