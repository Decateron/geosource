package transactions

import (
	"net/http"

	"github.com/joshheinrichs/httperr"
)

// CanModifyViewers returns true if the user with the given ID has permission
// to view viewers for the given channel, false otherwise.
func CanModifyViewers(userID, channelname string) (bool, httperr.Error) {
	return false, ErrNotImplemented
}

// CanViewViewers returns true if the user with the given ID has permission to
// modify viewers for the given channel, false otherwise.
func CanViewViewers(userID, channelname string) (bool, httperr.Error) {
	return false, ErrNotImplemented
}

// IsViewer returns true if the user with the given ID is viewer for the given
// channel, false otherwise.
func IsViewer(userID, channelname string) (bool, httperr.Error) {
	return false, ErrNotImplemented
}

// AddViewer adds the user with ID userID to the viewer list for the given
// channel. This transaction is executed under the permission level of the given
// requester. Returns an error if the requester does not have sufficient
// permission, or if some other error occurs within the database.
func AddViewer(requesterID, userID, channelname string) httperr.Error {
	permission, httpErr := CanModifyViewers(requesterID, channelname)
	if httpErr != nil {
		return httpErr
	} else if !permission {
		return ErrInsufficientPermission
	}
	err := db.Exec("INSERT INTO channel_viewers (chv_userid, chv_channelname) VALUES (?, ?);", userID, channelname).Error
	if err != nil {
		return httperr.New(err.Error(), http.StatusInternalServerError)
	}
	return nil
}

// GetViewers returns the list of users which are viewers for the given channel.
// This transaction is executed under the permission level of the given
// requester. Returns an error if the requester does not have sufficient
// permission, or if some other error occurs within the database.
func GetViewers(requesterID, channelname string) ([]string, httperr.Error) {
	permission, httpErr := CanViewViewers(requesterID, channelname)
	if httpErr != nil {
		return nil, httpErr
	} else if !permission {
		return nil, ErrInsufficientPermission
	}
	var viewers []string
	err := db.Table("channel_viewers").Select("chv_userid").Where("chv_channelname = ?", channelname).Scan(&viewers).Error
	if err != nil {
		return nil, httperr.New(err.Error(), http.StatusInternalServerError)
	}
	return viewers, nil
}

// RemoveViewer removes the user with ID userID from the viewer list for the
// given channel. This transaction is executed under the permission level of
// the given requester. Returns an error if the requester does not have
// sufficient permission, or if some other error occurs within the database.
func RemoveViewer(requesterID, userID, channelname string) httperr.Error {
	permission, httpErr := CanModifyViewers(requesterID, channelname)
	if httpErr != nil {
		return httpErr
	} else if !permission {
		return ErrInsufficientPermission
	}
	err := db.Exec("DELETE FROM channel_viewers WHERE chv_userid = ? and chv_channelname = ?", userID, channelname).Error
	if err != nil {
		return httperr.New(err.Error(), http.StatusInternalServerError)
	}
	return nil
}
