package transactions

import (
	"github.com/joshheinrichs/geosource/server/types"
)

func AddChannel(channel *types.Channel) error {
	return db.Create(channel).Error
}

func GetChannel(requesterUid, channelname string) (*types.Channel, error) {
	// TODO: Account for requester permission
	var channel types.Channel
	err := db.Where("ch_channelname = ?", channelname).First(&channel).Error
	return &channel, err
}

func GetChannels(requesterUid string) ([]string, error) {
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

func RemoveChannel(requesterUid, channelname string) error {
	return nil
}

func IsChannelCreator(requesterUid, uid, channelname string) (bool, error) {
	return false, nil
}
