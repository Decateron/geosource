package transactions

import (
	"github.com/joshheinrichs/geosource/server/types"
)

func AddComment(requesterUid string, comment *types.Comment) error {
	return nil
}

// func GetComments

func RemoveComment(requesterUid, commentId string) error {
	return nil
}

func IsCommentCreator(requesterUid, uid, commentId string) (bool, error) {
	return false, nil
}
