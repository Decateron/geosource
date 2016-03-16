package types

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	MAX_COMMENT_LENGTH = 500
)

type Comment struct {
	PostID    string    `json:"post" gorm:"column:cmt_postid"`
	CreatorID string    `json:"user" gorm:"column:cmt_userid_creator"`
	ID        string    `json:"id" gorm:"column:cmt_commentid"`
	ParentID  *string   `json:"parent" gorm:"column:cmt_commentid_parent"`
	Comment   string    `json:"comment" gorm:"column:cmt_comment"`
	Time      time.Time `json:"time" gorm:"column:cmt_time"`
}

func (comment *Comment) Validate() error {
	comment.Comment = strings.TrimSpace(comment.Comment)
	if len(comment.Comment) == 0 {
		return errors.New("Comment cannot be empty.")
	} else if len(comment.Comment) > MAX_COMMENT_LENGTH {
		return errors.New(fmt.Sprintf("Length of comment cannot exceed %i characters.", MAX_COMMENT_LENGTH))
	}
	return nil
}

func (comment *Comment) TableName() string {
	return "comments"
}
