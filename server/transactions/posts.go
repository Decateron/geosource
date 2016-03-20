package transactions

import (
	"errors"
	"time"

	"github.com/joshheinrichs/geosource/server/types"
)

type PostQuery struct {
	Flags struct {
		// If true, all posts will be searched, regardless of the channels and
		// other flags that were set in the query.
		All bool `url:"all"`
		// If true, posts that were created by the user will be included in
		// the search results.
		Mine bool `url:"mine"`
		// If true, posts that were favorited by the user will be included in
		// the search results.
		Favorites bool `url:"favorites"`
	} `url:"flags"`
	Time struct {
		Start time.Time `url:"start"`
		End   time.Time `url:"end"`
	} `url:"time"`
	Location struct {
		// The upper left bound of the region
		Start types.Location `url:"start"`
		// The lower right bound of the region
		End types.Location `url:"end"`
	}
	// The set of channels to query. If the all flag is true, this field is
	// ignored.
	Channels []string `url:"channels"`
}

func IsPostCreator(requester, userID, postID string) (bool, error) {
	return false, errors.New("function has not yet been implemented.")
}

func AddPost(requester string, post *types.Post) error {
	return db.Create(post).Error
}

func GetPosts(requester string) ([]*types.PersonalizedPostInfo, error) {
	var posts []*types.PersonalizedPostInfo
	err := db.Table("posts").
		Joins("LEFT JOIN user_favorites ON (p_postid = uf_postid)").
		Joins("LEFT JOIN users ON (u_userid = p_userid_creator)").
		Select("*, (uf_postid IS NOT NULL) AS favorited").
		Order("p_time desc").Find(&posts).Error
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
