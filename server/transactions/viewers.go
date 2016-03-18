package transactions

import "errors"

func CanModifyViewers(requester, channelname string) (bool, error) {
	return false, errors.New("function has not yet been implemented.")
}

func CanViewViewers(requester, channelname string) (bool, error) {
	return false, errors.New("function has not yet been implemented.")
}

func IsViewer(userID, channelname string) (bool, error) {
	return false, errors.New("function has not yet been implemented.")
}

func AddViewer(requester, userID, channelname string) error {
	permission, err := CanModifyViewers(requester, channelname)
	if err != nil {
		return err
	} else if !permission {
		return ErrInsufficientPermission
	}
	return db.Exec("INSERT INTO channel_viewers (chv_userid, chv_channelname) VALUES (?, ?);", userID, channelname).Error
}

func GetViewers(requester, channelname string) ([]string, error) {
	permission, err := CanViewViewers(requester, channelname)
	if err != nil {
		return nil, err
	} else if !permission {
		return nil, ErrInsufficientPermission
	}
	var viewers []string
	err = db.Table("channel_viewers").Select("chv_userid").Where("chv_channelname = ?", channelname).Scan(&viewers).Error
	if err != nil {
		return nil, err
	}
	return viewers, nil
}

func RemoveViewer(requester, userID, channelname string) error {
	permission, err := CanModifyViewers(requester, channelname)
	if err != nil {
		return err
	} else if !permission {
		return ErrInsufficientPermission
	}
	return db.Exec("DELETE FROM channel_viewers WHERE chv_userid = ? and chv_channelname = ?", userID, channelname).Error
}
