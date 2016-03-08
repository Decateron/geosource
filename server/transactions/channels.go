package transactions

import (
	"errors"

	"github.com/joshheinrichs/geosource/server/types"
)

func AddChannel(channel *types.Channel) error {
	return db.Create(channel).Error
}

func GetChannel(requester, channelname string) (*types.Channel, error) {
	// TODO: Account for requester permission
	var channel types.Channel
	err := db.Where("ch_channelname = ?", channelname).First(&channel).Error
	return &channel, err
}

func GetChannels(requester string) ([]string, error) {
	var channels []*types.Channel
	err := db.Order("ch_channelname").Find(&channels).Error
	if err != nil {
		return nil, err
	}
	channelnames := make([]string, len(channels))
	for i, channel := range channels {
		channelnames[i] = channel.Name
	}
	return channelnames, nil
}

func RemoveChannel(requester, channelname string) error {
	return errors.New("function has not yet been implemented.")
}

func IsChannelCreator(userid, channelname string) (bool, error) {
	return false, errors.New("function has not yet been implemented.")
}
