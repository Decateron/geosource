package transactions

// CanViewSubscriptions returns true if the requester can view the user with
// ID userID's subscriptions, false otherwise.
func CanViewSubscriptions(requesterID, userID string) (bool, error) {
	return requesterID == userID, nil
}

func CanModifySubscriptions(requesterID, userID string) (bool, error) {
	return requesterID == userID, nil
}

func AddSubscription(requesterID, userID, channelname string) error {
	permission, err := CanModifySubscriptions(requesterID, userID)
	if err != nil {
		return err
	} else if !permission {
		return ErrInsufficientPermission
	}
	return db.Exec("INSERT INTO user_subscriptions (us_userid, us_channelname) VALUES (?, ?)", requesterID, channelname).Error
}

func GetSubscriptions(requesterID, userID string) ([]string, error) {
	permission, err := CanViewSubscriptions(requesterID, userID)
	if err != nil {
		return nil, err
	} else if !permission {
		return nil, ErrInsufficientPermission
	}
	var subscriptions []struct {
		PostID string `gorm:"column:us_channelname"`
	}
	err = db.Table("user_subscriptions").
		Where("us_userid = ?", userID).
		Find(&subscriptions).Error
	if err != nil {
		return nil, err
	}
	channelnames := make([]string, len(subscriptions))
	for i, subscription := range subscriptions {
		channelnames[i] = subscription.PostID
	}
	return channelnames, nil
}

func RemoveSubscription(requesterID, userID, channelname string) error {
	permission, err := CanModifySubscriptions(requesterID, userID)
	if err != nil {
		return err
	} else if !permission {
		return ErrInsufficientPermission
	}
	return db.Exec("DELETE FROM user_subscriptions WHERE us_userid = ? AND us_channelname = ?", requesterID, channelname).Error
}
