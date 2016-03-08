package transactions

import "errors"

func CanViewModerators(requester, channelname string) (bool, error) {
	return false, errors.New("function has not yet been implemented.")
}

func CanModifyModerators(requester, channelname string) (bool, error) {
	return false, errors.New("function has not yet been implemented.")
}

func IsModerator(userid, channelname string) (bool, error) {
	return false, errors.New("function has not yet been implemented.")
}

func AddModerator(requester, userid, channelname string) error {
	permission, err := CanModifyModerators(requester, channelname)
	if err != nil {
		return err
	} else if !permission {
		return ErrInsufficientPermission
	}
	return db.Exec("INSERT INTO channel_moderators (chm_userid, chm_channelname) VALUES (?, ?);", userid, channelname).Error
}

func GetModerators(requester, channelname string) ([]string, error) {
	permission, err := CanViewModerators(requester, channelname)
	if err != nil {
		return nil, err
	} else if !permission {
		return nil, ErrInsufficientPermission
	}
	var moderators []string
	err = db.Table("channel_moderators").Select("chv_userid").Where("chv_channelname = ?", channelname).Scan(&moderators).Error
	if err != nil {
		return nil, err
	}
	return moderators, nil
}

func RemoveModerator(requester, userid, channelname string) error {
	permission, err := CanModifyModerators(requester, channelname)
	if err != nil {
		return err
	} else if !permission {
		return ErrInsufficientPermission
	}
	return db.Exec("DELETE FROM channel_moderators WHERE chm_userid = ? and chm_channelname = ?", userid, channelname).Error
}
