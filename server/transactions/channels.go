package transactions

import (
	"errors"

	"github.com/joshheinrichs/geosource/server/types"
)

func IsChannelCreator(userID, channelname string) (bool, error) {
	return false, errors.New("function has not yet been implemented.")
}

func AddChannel(channel *types.Channel) error {
	return db.Create(channel).Error
}

func GetChannel(requester, channelname string) (*types.PersonalizedChannel, error) {
	// TODO: Account for requester permission
	var channel types.PersonalizedChannel
	err := db.Table("channels").
		Joins("LEFT JOIN user_subscriptions ON (ch_channelname = us_channelname AND us_userid = ?)", requester).
		Joins("LEFT JOIN users ON (u_userid = ch_userid_creator)").
		Select("*, (us_channelname IS NOT NULL) AS subscribed").
		First(&channel).Error
	return &channel, err
}

func GetChannels(requester string) ([]*types.PersonalizedChannelInfo, error) {
	var channels []*types.PersonalizedChannelInfo
	err := db.Table("channels").
		Joins("LEFT JOIN user_subscriptions ON (ch_channelname = us_channelname AND us_userid = ?)", requester).
		Joins("LEFT JOIN users ON (u_userid = ch_userid_creator)").
		Select("*, (us_channelname IS NOT NULL) AS subscribed").
		Order("ch_channelname").Find(&channels).Error
	if err != nil {
		return nil, err
	}
	return channels, nil
}

func RemoveChannel(requester, channelname string) error {
	return errors.New("function has not yet been implemented.")
}
