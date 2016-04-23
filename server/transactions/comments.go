package transactions

import (
	"github.com/jinzhu/gorm"
	"github.com/joshheinrichs/geosource/server/types"
)

func IsCommentCreator(requesterID, userID, commentID string) (bool, error) {
	var comment types.Comment
	err := db.Where("cmt_commentid = ?", commentID).First(&comment).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

func AddComment(requesterID string, comment *types.Comment) error {
	return db.Create(comment).Error
}

// GetComment returns a comment with the given ID, or nil if it does not exist.
// An error is returned if some issue occurred with the database.
func GetComment(requesterID, commentID string) (*types.Comment, error) {
	var comment types.Comment
	err := db.Where("cmt_commentid = ?", commentID).First(&comment).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &comment, nil
}

func GetComments(requesterID, postID string) ([]*types.Comment, error) {
	var comments []*types.Comment
	err := db.Where("cmt_postid = ?", postID).Order("cmt_time").Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func RemoveComment(requesterID, commentID string) error {
	return db.Where("cmd_commentid = ?", commentID).Delete(&types.Comment{}).Error
}
