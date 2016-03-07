package transactions

import (
	"errors"

	"github.com/joshheinrichs/geosource/server/types"
)

func AddComment(requesterUid string, comment *types.Comment) error {
	return errors.New("function has not yet been implemented.")
}

// func GetComments

func RemoveComment(requesterUid, commentId string) error {
	return errors.New("function has not yet been implemented.")
}

func IsCommentCreator(requesterUid, uid, commentId string) (bool, error) {
	return false, errors.New("function has not yet been implemented.")
}
