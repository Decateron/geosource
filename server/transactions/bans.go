package transactions

import (
	"errors"
)

// CanViewBans returns true if the user with the given ID has permission to view
// bans for the given channel, false otherwise.
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

// CanModifyBans returns true if the user with the given ID has permission to
// modify bans for the given channel, false otherwise.
func CanModifyBans(userID, channelname string) (bool, error) {
	return false, errors.New("function has not yet been implemented")
}

// IsBanned returns true if the user with the given ID is banned from the given
// channel, false otherwise.
func IsBanned(userID, channelname string) (bool, error) {
	return false, errors.New("function has not yet been implemented")
}

// AddBan adds the user with ID userID to the ban list for the given channel.
// This transaction is executed under the permission level of the given
// requester. Returns an error if the requester does not have sufficient
// permission, or if some other error occurs within the database.
func AddBan(requesterID, userID, channelname string) error {
	permission, err := CanModifyBans(requesterID, channelname)
	if err != nil {
		return err
	} else if !permission {
		return ErrInsufficientPermission
	}
	return db.Exec("INSERT INTO channel_bans (chb_userid, chb_channelname) VALUES (?, ?);", userID, channelname).Error
}

// GetBans returns the list of users which are banned from the given channel.
// This transaction is executed under the permission level of the given
// requester. Returns an error if the requester does not have sufficient
// permission, or if some other error occurs within the database.
func GetBans(requesterID, channelname string) ([]string, error) {
	permission, err := CanViewBans(requesterID, channelname)
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

// RemoveBan removes the user with ID userID from the ban list for the given
// channel. This transaction is executed under the permission level of the given
// requester. Returns an error if the requester does not have sufficient
// permission, or if some other error occurs within the database.
func RemoveBan(requesterID, userID, channelname string) error {
	permission, err := CanModifyBans(requesterID, channelname)
	if err != nil {
		return err
	} else if !permission {
		return ErrInsufficientPermission
	}
	return db.Exec("DELETE FROM channel_bans WHERE chb_userid = ? and chb_channelname = ?", userID, channelname).Error
}
