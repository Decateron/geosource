package transactions

// CanViewSubscriptions returns true if the requester can view the user with
// ID userID's subscriptions, false otherwise.
func CanViewSubscriptions(requester, userID string) (bool, error) {
	return requester == userID, nil
}

func CanModifySubscriptions(requester, userID string) (bool, error) {
	return requester == userID, nil
}

func AddSubscription(requester, userID, channelname string) error {
	permission, err := CanModifySubscriptions(requester, userID)
	if err != nil {
		return err
	} else if !permission {
		return ErrInsufficientPermission
	}
	return db.Exec("INSERT INTO user_subscriptions (us_userid, us_channelname) VALUES (?, ?)", requester, channelname).Error
}

func GetSubscriptions(requester, userID string) ([]string, error) {
	permission, err := CanViewSubscriptions(requester, userID)
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

func RemoveSubscription(requester, userID, channelname string) error {
	permission, err := CanModifySubscriptions(requester, userID)
	if err != nil {
		return err
	} else if !permission {
		return ErrInsufficientPermission
	}
	return db.Exec("DELETE FROM user_subscriptions WHERE us_userid = ? AND us_channelname = ?", requester, channelname).Error
}
