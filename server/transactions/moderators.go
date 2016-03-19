package transactions

import "errors"

// Returns true if the user with the given ID has permission to view moderators
// for the given channel, false otherwise.
func CanViewModerators(userID, channelname string) (bool, error) {
	return false, errors.New("function has not yet been implemented.")
}

// Returns true if the user with the given ID has permission to modify
// moderators for the given channel, false otherwise.
func CanModifyModerators(userID, channelname string) (bool, error) {
	return false, errors.New("function has not yet been implemented.")
}

// Returns true if the user with the given ID is moderator for the given
// channel, false otherwise.
func IsModerator(userID, channelname string) (bool, error) {
	return false, errors.New("function has not yet been implemented.")
}

// Adds the user with ID userID to the moderator list for the given channel.
// This transaction is executed under the permission level of the given
// requester. Returns an error if the requester does not have sufficient
// permission, or if some other error occurs within the database.
func AddModerator(requester, userID, channelname string) error {
	permission, err := CanModifyModerators(requester, channelname)
	if err != nil {
		return err
	} else if !permission {
		return ErrInsufficientPermission
	}
	return db.Exec("INSERT INTO channel_moderators (chm_userid, chm_channelname) VALUES (?, ?);", userID, channelname).Error
}

// Returns the list of users which are moderators for the given channel. This
// transaction is executed under the permission level of the given requester.
// Returns an error if the requester does not have sufficient permission, or
// if some other error occurs within the database.
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

// Removes the user with ID userID from the moderator list for the given
// channel. This transaction is executed under the permission level of the
// given requester. Returns an error if the requester does not have sufficient
// permission, or if some other error occurs within the database.
func RemoveModerator(requester, userID, channelname string) error {
	permission, err := CanModifyModerators(requester, channelname)
	if err != nil {
		return err
	} else if !permission {
		return ErrInsufficientPermission
	}
	return db.Exec("DELETE FROM channel_moderators WHERE chm_userid = ? and chm_channelname = ?", userID, channelname).Error
}
