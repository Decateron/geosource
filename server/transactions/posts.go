package transactions

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/joshheinrichs/geosource/server/types"
)

type PostQueryParams struct {
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
	TimeRange *struct {
		Start time.Time `url:"start"`
		End   time.Time `url:"end"`
	} `url:"timeRange"`
	LocationRange *struct {
		// The upper left bound of the region
		Min types.Location `url:"min"`
		// The lower right bound of the region
		Max types.Location `url:"max"`
	} `url:"locationRange"`
	// The set of channels to query. If the all flag is true, this field is
	// ignored.
	Channels []string `url:"channels"`
}

func IsPostCreator(requester, userID, postID string) (bool, error) {
	return false, errors.New("function has not yet been implemented.")
}

func AddPost(requester string, post *types.Post) error {
	jsonFields, err := json.Marshal(post.Fields)
	if err != nil {
		return err
	}
	return db.Exec("INSERT INTO posts (p_postid, p_userid_creator, p_channelname, p_title, p_thumbnail, p_time, p_location, p_fields) VALUES (?, ?, ?, ?, ?, ?, ST_MakePoint(?,?), ?)",
		post.ID, post.CreatorID, post.Channel, post.Title, post.Thumbnail, post.Time,
		post.Location.Longitude, post.Location.Latitude, jsonFields).Error
	// return db.Create(post).Error
}

func GetPosts(requester string) ([]*types.PersonalizedPostInfo, error) {
	var posts []*types.PersonalizedPostInfo
	err := db.Table("posts").
		Joins("LEFT JOIN user_favorites ON (p_postid = uf_postid AND uf_userid = ?)", requester).
		Joins("LEFT JOIN users ON (u_userid = p_userid_creator)").
		Select("*, (uf_postid IS NOT NULL) AS favorited, ST_AsText(p_location) AS location").
		Order("p_time desc").Find(&posts).Error

	// if postQueryParams.TimeRange != nil {
	// 	query.Where("p_time > ? AND p_time < ?",
	// 		postQueryParams.TimeRange.Start,
	// 		postQueryParams.TimeRange.End)
	// }
	// if postQueryParams.LocationRange != nil {
	// 	query.Where("p_location && ST_MakeEnvelope (?, ?, ?, ?, 4326)",
	// 		postQueryParams.LocationRange.Min.Longitude,
	// 		postQueryParams.LocationRange.Min.Latitude,
	// 		postQueryParams.LocationRange.Max.Longitude,
	// 		postQueryParams.LocationRnage.Max.Latitude)
	// }
	// err := query.Select("*, (uf_postid IS NOT NULL) AS favorited").
	// Order("p_time desc").Find(&posts).Error

	if err != nil {
		return nil, err
	}
	return posts, nil
}

func GetPost(requester, postID string) (*types.PersonalizedPost, error) {
	var post types.PersonalizedPost
	err := db.Table("posts").
		Where("p_postid = ?", postID).
		Joins("LEFT JOIN user_favorites ON (p_postid = uf_postid AND uf_userid = ?)", requester).
		Joins("LEFT JOIN users ON (u_userid = p_userid_creator)").
		Select("*, (uf_postid IS NOT NULL) AS favorited, ST_AsText(p_location) AS location").
		First(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func RemovePost(requester, postID string) error {
	return errors.New("function has not yet been implemented.")
}
