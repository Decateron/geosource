package transactions

import (
	"errors"
)

func CanModifyModerators(requester, channelname string) (bool, error) {
	return false, errors.New("function has not yet been implemented.")
}

func CanModifyViewers(requester, channelname string) (bool, error) {
	return false, errors.New("function has not yet been implemented.")
}

func CanModifyBans(requester, channelname string) (bool, error) {
	return false, errors.New("function has not yet been implemented.")
}
