package transactions

import (
	"errors"
)

func CanViewBans(userid, channelname string) (bool, error) {
	permission, err := IsAdmin(userid)
	if err != nil {
		return false, err
	} else if permission {
		return true, nil
	}
	permission, err = IsChannelCreator(userid, channelname)
	if err != nil {
		return false, err
	} else if permission {
		return true, nil
	}
	permission, err = IsModerator(userid, channelname)
	if err != nil {
		return false, err
	} else if permission {
		return true, nil
	}
	return false, nil
}

func CanModifyBans(userid, channelname string) (bool, error) {
	return false, errors.New("function has not yet been implemented.")
}

func IsBanned(userid, channelname string) (bool, error) {
	return false, errors.New("function has not yet been implemented.")
}

func AddBan(requester, userid, channelname string) error {
	permission, err := CanModifyBans(requester, channelname)
	if err != nil {
		return err
	} else if !permission {
		return ErrInsufficientPermission
	}
	return db.Exec("INSERT INTO channel_bans (chb_user, chb_channelname) VALUES (?, ?);", userid, channelname).Error
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

func RemoveBan(requester, userid, channelname string) error {
	permission, err := CanModifyBans(requester, channelname)
	if err != nil {
		return err
	} else if !permission {
		return ErrInsufficientPermission
	}
	return db.Exec("DELETE FROM channel_bans WHERE chb_userid = ? and chb_channelname = ?", userid, channelname).Error
}
