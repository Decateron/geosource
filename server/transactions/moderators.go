package transactions

import (
	"net/http"

	"github.com/joshheinrichs/httperr"
)

// CanViewModerators returns true if the user with the given ID has permission
// to view moderators for the given channel, false otherwise.
func CanViewModerators(userID, channelname string) (bool, httperr.Error) {
	return false, ErrNotImplemented
}

// CanModifyModerators returns true if the user with the given ID has permission
// to modify moderators for the given channel, false otherwise.
func CanModifyModerators(userID, channelname string) (bool, httperr.Error) {
	return false, ErrNotImplemented
}

// IsModerator returns true if the user with the given ID is moderator for the
// given channel, false otherwise.
func IsModerator(userID, channelname string) (bool, httperr.Error) {
	return false, ErrNotImplemented
}

// AddModerator adds the user with ID userID to the moderator list for the given
// channel. This transaction is executed under the permission level of the given
// requester. Returns an error if the requester does not have sufficient
// permission, or if some other error occurs within the database.
func AddModerator(requesterID, userID, channelname string) httperr.Error {
	permission, httpErr := CanModifyModerators(requesterID, channelname)
	if httpErr != nil {
		return httpErr
	} else if !permission {
		return ErrInsufficientPermission
	}
	err := db.Exec("INSERT INTO channel_moderators (chm_userid, chm_channelname) VALUES (?, ?);", userID, channelname).Error
	if err != nil {
		return httperr.New(err.Error(), http.StatusInternalServerError)
	}
	return nil
}

// GetModerators returns the list of users which are moderators for the given
// channel. This transaction is executed under the permission level of the given
// requester. Returns an error if the requester does not have sufficient
// permission, or if some other error occurs within the database.
func GetModerators(requesterID, channelname string) ([]string, httperr.Error) {
	permission, httpErr := CanViewModerators(requesterID, channelname)
	if httpErr != nil {
		return nil, httpErr
	} else if !permission {
		return nil, ErrInsufficientPermission
	}
	var moderators []string
	err := db.Table("channel_moderators").Select("chv_userid").Where("chv_channelname = ?", channelname).Scan(&moderators).Error
	if err != nil {
		return nil, httperr.New(err.Error(), http.StatusInternalServerError)
	}
	return moderators, nil
}

// RemoveModerator removes the user with ID userID from the moderator list for
// the given channel. This transaction is executed under the permission level of
// the given requester. Returns an error if the requester does not have
// sufficient permission, or if some other error occurs within the database.
func RemoveModerator(requesterID, userID, channelname string) httperr.Error {
	permission, httpErr := CanModifyModerators(requesterID, channelname)
	if httpErr != nil {
		return httpErr
	} else if !permission {
		return ErrInsufficientPermission
	}
	err := db.Exec("DELETE FROM channel_moderators WHERE chm_userid = ? and chm_channelname = ?", userID, channelname).Error
	if err != nil {
		return httperr.New(err.Error(), http.StatusInternalServerError)
	}
	return nil
}
