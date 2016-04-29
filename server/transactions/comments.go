package transactions

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/joshheinrichs/geosource/server/types"
	"github.com/joshheinrichs/httperr"
)

func IsCommentCreator(requesterID, userID, commentID string) (bool, httperr.Error) {
	var comment types.Comment
	err := db.Where("cmt_commentid = ?", commentID).First(&comment).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	} else if err != nil {
		return false, httperr.New(err.Error(), http.StatusInternalServerError)
	}
	return true, nil
}

func AddComment(requesterID string, comment *types.Comment) httperr.Error {
	err := db.Create(comment).Error
	if err != nil {
		return httperr.New(err.Error(), http.StatusInternalServerError)
	}
	return nil
}

// GetComment returns a comment with the given ID, or nil if it does not exist.
// An error is returned if some issue occurred with the database.
func GetComment(requesterID, commentID string) (*types.Comment, httperr.Error) {
	var comment types.Comment
	err := db.Where("cmt_commentid = ?", commentID).First(&comment).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, httperr.New(err.Error(), http.StatusInternalServerError)
	}
	return &comment, nil
}

func GetComments(requesterID, postID string) ([]*types.Comment, httperr.Error) {
	var comments []*types.Comment
	err := db.Where("cmt_postid = ?", postID).Order("cmt_time").Find(&comments).Error
	if err != nil {
		return nil, httperr.New(err.Error(), http.StatusInternalServerError)
	}
	return comments, nil
}

func RemoveComment(requesterID, commentID string) httperr.Error {
	err := db.Where("cmd_commentid = ?", commentID).Delete(&types.Comment{}).Error
	if err != nil {
		return httperr.New(err.Error(), http.StatusInternalServerError)
	}
	return nil
}
