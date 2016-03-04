package types

import (
	"time"

	"github.com/joshheinrichs/geosource/server/types/fields"
)

// "github.com/joshheinrichs/geosource/server/transactions"

type PostInfo struct {
	Id        string    `json:"id" gorm:"column:p_postid"`
	CreatorId string    `json:"creator" gorm:"column:p_userid_creator"`
	Channel   string    `json:"channel" gorm:"column:p_channelname"`
	Title     string    `json:"title" gorm:"column:p_title"`
	Time      time.Time `json:"time" gorm:"column:p_time"`
	Location  Location  `json:"location" gorm:"column:p_location" sql:"type:POINT NOT NULL"`
}

func (postInfo *PostInfo) TableName() string {
	return "posts"
}

type Post struct {
	PostInfo
	Fields fields.Fields `json:"fields" gorm:"column:p_fields" sql:"type:JSONB NOT NULL"`
}

func (post *Post) TableName() string {
	return "posts"
}

type Submission struct {
	Title    string         `json:"title"`
	Channel  string         `json:"channel"`
	Location Location       `json:"location"`
	Values   []fields.Value `json:"values"`
}
