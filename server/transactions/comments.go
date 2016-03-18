package transactions

import (
	"github.com/jinzhu/gorm"
	"github.com/joshheinrichs/geosource/server/types"
)

func IsCommentCreator(requester, userID, commentID string) (bool, error) {
	var comment types.Comment
	err := db.Where("cmt_commentid = ?", commentID).First(&comment).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

func AddComment(requester string, comment *types.Comment) error {
	return db.Create(comment).Error
}

func GetComments(requester, postID string) ([]*types.Comment, error) {
	var comments []*types.Comment
	err := db.Where("cmt_postid = ?", postID).Order("cmt_time").Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func RemoveComment(requester, commentID string) error {
	return db.Where("cmd_commentid = ?", commentID).Delete(&types.Comment{}).Error
}
