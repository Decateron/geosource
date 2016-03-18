package transactions

import (
	"errors"
)

func CanViewBans(userID, channelname string) (bool, error) {
	permission, err := IsAdmin(userID)
	if err != nil {
		return false, err
	} else if permission {
		return true, nil
	}
	permission, err = IsChannelCreator(userID, channelname)
	if err != nil {
		return false, err
	} else if permission {
		return true, nil
	}
	permission, err = IsModerator(userID, channelname)
	if err != nil {
		return false, err
	} else if permission {
		return true, nil
	}
	return false, nil
}

func CanModifyBans(userID, channelname string) (bool, error) {
	return false, errors.New("function has not yet been implemented.")
}

func IsBanned(userID, channelname string) (bool, error) {
	return false, errors.New("function has not yet been implemented.")
}

func AddBan(requester, userID, channelname string) error {
	permission, err := CanModifyBans(requester, channelname)
	if err != nil {
		return err
	} else if !permission {
		return ErrInsufficientPermission
	}
	return db.Exec("INSERT INTO channel_bans (chb_userid, chb_channelname) VALUES (?, ?);", userID, channelname).Error
}

func GetBans(requester, channelname string) ([]string, error) {
	permission, err := CanViewBans(requester, channelname)
	if err != nil {
		return nil, err
	} else if !permission {
		return nil, ErrInsufficientPermission
	}
	var bans []string
	err = db.Table("channel_bans").Select("chb_userid").Where("chb_channelname = ?", channelname).Scan(&bans).Error
	if err != nil {
		return nil, err
	}
	return bans, nil
}

func RemoveBan(requester, userID, channelname string) error {
	permission, err := CanModifyBans(requester, channelname)
	if err != nil {
		return err
	} else if !permission {
		return ErrInsufficientPermission
	}
	return db.Exec("DELETE FROM channel_bans WHERE chb_userid = ? and chb_channelname = ?", userID, channelname).Error
}
