package transactions

import (
	"net/http"

	"github.com/joshheinrichs/httperr"
)

// CanViewSubscriptions returns true if the requester can view the user with
// ID userID's subscriptions, false otherwise.
func CanViewSubscriptions(requesterID, userID string) (bool, httperr.Error) {
	return requesterID == userID, nil
}

func CanModifySubscriptions(requesterID, userID string) (bool, httperr.Error) {
	return requesterID == userID, nil
}

func AddSubscription(requesterID, userID, channelname string) httperr.Error {
	permission, httpErr := CanModifySubscriptions(requesterID, userID)
	if httpErr != nil {
		return httpErr
	} else if !permission {
		return ErrInsufficientPermission
	}
	err := db.Exec("INSERT INTO user_subscriptions (us_userid, us_channelname) VALUES (?, ?)", requesterID, channelname).Error
	if err != nil {
		return httperr.New(err.Error(), http.StatusInternalServerError)
	}
	return nil
}

func GetSubscriptions(requesterID, userID string) ([]string, httperr.Error) {
	permission, httpErr := CanViewSubscriptions(requesterID, userID)
	if httpErr != nil {
		return nil, httpErr
	} else if !permission {
		return nil, ErrInsufficientPermission
	}
	var subscriptions []struct {
		PostID string `gorm:"column:us_channelname"`
	}
	err := db.Table("user_subscriptions").
		Where("us_userid = ?", userID).
		Find(&subscriptions).Error
	if err != nil {
		return nil, httperr.New(err.Error(), http.StatusInternalServerError)
	}
	channelnames := make([]string, len(subscriptions))
	for i, subscription := range subscriptions {
		channelnames[i] = subscription.PostID
	}
	return channelnames, nil
}

func RemoveSubscription(requesterID, userID, channelname string) httperr.Error {
	permission, httpErr := CanModifySubscriptions(requesterID, userID)
	if httpErr != nil {
		return httpErr
	} else if !permission {
		return ErrInsufficientPermission
	}
	err := db.Exec("DELETE FROM user_subscriptions WHERE us_userid = ? AND us_channelname = ?", requesterID, channelname).Error
	if err != nil {
		return httperr.New(err.Error(), http.StatusInternalServerError)
	}
	return nil
}
