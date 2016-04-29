package transactions

import (
	"net/http"

	"github.com/joshheinrichs/geosource/server/types"
	"github.com/joshheinrichs/httperr"
)

func IsChannelCreator(userID, channelname string) (bool, httperr.Error) {
	return false, ErrNotImplemented
}

func AddChannel(channel *types.Channel) httperr.Error {
	err := db.Create(channel).Error
	if err != nil {
		return httperr.New(err.Error(), http.StatusInternalServerError)
	}
	return nil
}

func GetChannel(requesterID, channelname string) (*types.PersonalizedChannel, httperr.Error) {
	// TODO: Account for requester permission
	var channel types.PersonalizedChannel
	err := db.Table("channels").
		Joins("LEFT JOIN user_subscriptions ON (ch_channelname = us_channelname AND us_userid = ?)", requesterID).
		Joins("LEFT JOIN users ON (u_userid = ch_userid_creator)").
		Select("*, (us_channelname IS NOT NULL) AS subscribed").
		Where("ch_channelname = ?", channelname).
		First(&channel).Error
	if err != nil {
		return nil, httperr.New(err.Error(), http.StatusInternalServerError)
	}
	return &channel, nil
}

func GetChannels(requesterID string) ([]*types.PersonalizedChannelInfo, httperr.Error) {
	var channels []*types.PersonalizedChannelInfo
	err := db.Table("channels").
		Joins("LEFT JOIN user_subscriptions ON (ch_channelname = us_channelname AND us_userid = ?)", requesterID).
		Joins("LEFT JOIN users ON (u_userid = ch_userid_creator)").
		Select("*, (us_channelname IS NOT NULL) AS subscribed").
		Order("ch_channelname").Find(&channels).Error
	if err != nil {
		return nil, httperr.New(err.Error(), http.StatusInternalServerError)
	}
	return channels, nil
}

func RemoveChannel(requesterID, channelname string) httperr.Error {
	return ErrNotImplemented
}
