package transactions

import (
	"errors"

	"github.com/joshheinrichs/geosource/server/types"
)

func AddPost(requester string, post *types.Post) error {
	return db.Create(post).Error
}

func GetPosts(requester string) ([]*types.PostInfo, error) {
	var posts []*types.PostInfo
	err := db.Order("p_time desc").Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func GetPost(requester, postID string) (*types.Post, error) {
	var post types.Post
	err := db.Where("p_postid = ?", postID).First(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func RemovePost(requester, postID string) error {
	return errors.New("function has not yet been implemented.")
}

func IsPostCreator(requester, userID, postID string) (bool, error) {
	return false, errors.New("function has not yet been implemented.")
}
