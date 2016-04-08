package transactions

import "errors"

// CanModifyViewers returns true if the user with the given ID has permission
// to view viewers for the given channel, false otherwise.
func CanModifyViewers(userID, channelname string) (bool, error) {
	return false, errors.New("function has not yet been implemented")
}

// CanViewViewers returns true if the user with the given ID has permission to
// modify viewers for the given channel, false otherwise.
func CanViewViewers(userID, channelname string) (bool, error) {
	return false, errors.New("function has not yet been implemented")
}

// IsViewer returns true if the user with the given ID is viewer for the given
// channel, false otherwise.
func IsViewer(userID, channelname string) (bool, error) {
	return false, errors.New("function has not yet been implemented")
}

// AddViewer adds the user with ID userID to the viewer list for the given
// channel. This transaction is executed under the permission level of the given
// requester. Returns an error if the requester does not have sufficient
// permission, or if some other error occurs within the database.
func AddViewer(requester, userID, channelname string) error {
	permission, err := CanModifyViewers(requester, channelname)
	if err != nil {
		return err
	} else if !permission {
		return ErrInsufficientPermission
	}
	return db.Exec("INSERT INTO channel_viewers (chv_userid, chv_channelname) VALUES (?, ?);", userID, channelname).Error
}

// GetViewers returns the list of users which are viewers for the given channel.
// This transaction is executed under the permission level of the given
// requester. Returns an error if the requester does not have sufficient
// permission, or if some other error occurs within the database.
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

// RemoveViewer removes the user with ID userID from the viewer list for the
// given channel. This transaction is executed under the permission level of
// the given requester. Returns an error if the requester does not have
// sufficient permission, or if some other error occurs within the database.
func RemoveViewer(requester, userID, channelname string) error {
	permission, err := CanModifyViewers(requester, channelname)
	if err != nil {
		return err
	} else if !permission {
		return ErrInsufficientPermission
	}
	return db.Exec("DELETE FROM channel_viewers WHERE chv_userid = ? and chv_channelname = ?", userID, channelname).Error
}
