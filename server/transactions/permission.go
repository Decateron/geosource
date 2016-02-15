package transactions

func CanModifyModerators(requester, channelname string) (bool, error) {
	return false, nil
}

func CanModifyViewers(requester, channelname string) (bool, error) {
	return false, nil
}

func CanModifyBans(requester, channelname string) (bool, error) {
	return false, nil
}
