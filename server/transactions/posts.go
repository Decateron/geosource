package transactions

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/joshheinrichs/geosource/server/types"
)

type PostQueryParams struct {
	Flags         *Flags         `url:"flags"`
	TimeRange     *TimeRange     `url:"timeRange"`
	LocationRange *LocationRange `url:"locationRange"`
}

type Flags struct {
	// If true, all posts will be searched, regardless of the channels and
	// other flags that were set in the query.
	All bool `url:"all"`
	// If true, posts that were created by the user will be included in
	// the search results.
	Mine bool `url:"mine"`
	// If true, posts that were favorited by the user will be included in
	// the search results.
	Favorites bool `url:"favorites"`
}

type TimeRange struct {
	Min time.Time `url:"min"`
	Max time.Time `url:"max"`
}

type LocationRange struct {
	// The upper left bound of the region
	Min types.Location `url:"min"`
	// The lower right bound of the region
	Max types.Location `url:"max"`
}

func IsPostCreator(requester, userID, postID string) (bool, error) {
	return false, errors.New("function has not yet been implemented.")
}

func AddPost(requester string, post *types.Post) error {
	jsonFields, err := json.Marshal(post.Fields)
	if err != nil {
		return err
	}
	return db.Exec("INSERT INTO posts "+
		"(p_postid, p_userid_creator, p_channelname, p_title, p_thumbnail, p_time, p_location, p_fields) "+
		"VALUES (?, ?, ?, ?, ?, ?, ST_MakePoint(?,?), ?)",
		post.ID, post.CreatorID, post.Channel, post.Title, post.Thumbnail, post.Time,
		post.Location.Longitude, post.Location.Latitude, jsonFields).Error
}

func GetPosts(requester string, postQueryParams *PostQueryParams) ([]*types.PersonalizedPostInfo, error) {
	var posts []*types.PersonalizedPostInfo
	query := db.Table("posts").
		Joins("LEFT JOIN user_favorites ON (p_postid = uf_postid AND uf_userid = ?)", requester).
		Joins("LEFT JOIN users ON (u_userid = p_userid_creator)")

	if postQueryParams.Flags != nil {
		// TODO
	}
	if postQueryParams.LocationRange != nil {
		query = query.Where("p_location && ST_MakeEnvelope (?,?,?,?)",
			postQueryParams.LocationRange.Min.Longitude,
			postQueryParams.LocationRange.Min.Latitude,
			postQueryParams.LocationRange.Max.Longitude,
			postQueryParams.LocationRange.Max.Latitude)
	}
	if postQueryParams.TimeRange != nil {
		query = query.Where("p_time > ? AND p_time < ?",
			postQueryParams.TimeRange.Min,
			postQueryParams.TimeRange.Max)
	}

	err := query.Select("*, (uf_postid IS NOT NULL) AS favorited, ST_AsText(p_location) AS location").
		Order("p_time desc").Find(&posts).Error
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
