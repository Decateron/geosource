package transactions

import (
	"github.com/joshheinrichs/geosource/server/types"
)

func AddPost(requesterUid string, post *types.Post) error {
	return db.Create(post).Error
}

func GetPosts(requesterUid string) ([]*types.PostInfo, error) {
	var posts []*types.PostInfo
	err := db.Order("p_time desc").Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func GetPost(requesterUid, pid string) (*types.Post, error) {
	var post types.Post
	err := db.Where("p_postid = ?", pid).First(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func RemovePost(requesterUid, postId string) error {
	return nil
}

func IsPostCreator(requesterUid, uid, postId string) (bool, error) {
	return false, nil
}
