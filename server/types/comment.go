package types

import "time"

type Comment struct {
	PostId    string    `json:"post" gorm:"column:cmt_postid"`
	CreatorId string    `json:"user" gorm:"column:cmt_userid_creator"`
	Id        string    `json:"id" gorm:"column:cmt_commentid"`
	ParentId  *string   `json:"parent" gorm:"column:cmt_commentid_parent"`
	Comment   string    `json:"comment" gorm:"column:cmt_comment"`
	Time      time.Time `json:"time" gorm:"column:cmt_time"`
}

func (comment *Comment) TableName() string {
	return "comments"
}
