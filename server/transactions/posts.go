package transactions

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/joshheinrichs/geosource/server/types"
	"github.com/joshheinrichs/httperr"
)

type PostQueryParams struct {
	Flags         *Flags         `url:"flags"`
	TimeRange     *TimeRange     `url:"timeRange"`
	LocationRange *LocationRange `url:"locationRange"`
	Limit         *int           `url:"limit"`
	Offset        *int           `url:"offset"`
}

type Flags struct {
	// If true, posts that were created by the user will be included in
	// the search results.
	Mine bool `url:"mine"`
	// If true, posts that were favorited by the user will be included in
	// the search results.
	Favorites bool `url:"favorites"`
	// If true, posts that are in the user's subscribed channels will be
	// included in the search results.
	Subscriptions bool `url:"subscriptions"`
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

func IsPostCreator(requesterID, userID, postID string) (bool, httperr.Error) {
	return false, ErrNotImplemented
}

func AddPost(requesterID string, post *types.Post) httperr.Error {
	jsonFields, err := json.Marshal(post.Fields)
	if err != nil {
		return httperr.New(err.Error(), http.StatusInternalServerError)
	}
	err = db.Exec("INSERT INTO posts "+
		"(p_postid, p_userid_creator, p_channelname, p_title, p_thumbnail, p_time, p_location, p_fields) "+
		"VALUES (?, ?, ?, ?, ?, ?, ST_MakePoint(?,?), ?)",
		post.ID, post.CreatorID, post.Channel, post.Title, post.Thumbnail, post.Time,
		post.Location.Longitude, post.Location.Latitude, jsonFields).Error
	if err != nil {
		return httperr.New(err.Error(), http.StatusInternalServerError)
	}
	return nil
}

func GetPosts(requesterID string, postQueryParams *PostQueryParams) ([]*types.PersonalizedPostInfo, httperr.Error) {
	var posts []*types.PersonalizedPostInfo
	query := db.Table("posts").
		Joins("LEFT JOIN user_favorites ON (p_postid = uf_postid AND uf_userid = ?)", requesterID).
		Joins("LEFT JOIN user_subscriptions ON (p_channelname = us_channelname AND us_userid = ?)", requesterID).
		Joins("LEFT JOIN users ON (u_userid = p_userid_creator)").
		Select("*, (uf_postid IS NOT NULL) AS favorited, ST_AsText(p_location) AS location, (us_channelname IS NOT NULL) AS subscribed").
		Order("p_time desc")

	if postQueryParams.Flags != nil {
		if postQueryParams.Flags.Mine {
			query = query.Where("p_userid_creator = ?", requesterID)
		}
		if postQueryParams.Flags.Favorites {
			query = query.Where("uf_postid IS NOT NULL")
		}
		if postQueryParams.Flags.Subscriptions {
			query = query.Where("us_channelname IS NOT NULL")
		}
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
	if postQueryParams.Limit != nil {
		query = query.Limit(*postQueryParams.Limit)
	}
	if postQueryParams.Offset != nil {
		query = query.Offset(*postQueryParams.Offset)
	}

	err := query.Find(&posts).Error
	if err != nil {
		return nil, httperr.New(err.Error(), http.StatusInternalServerError)
	}
	return posts, nil
}

func GetPost(requesterID, postID string) (*types.PersonalizedPost, httperr.Error) {
	var post types.PersonalizedPost
	err := db.Table("posts").
		Where("p_postid = ?", postID).
		Joins("LEFT JOIN user_favorites ON (p_postid = uf_postid AND uf_userid = ?)", requesterID).
		Joins("LEFT JOIN users ON (u_userid = p_userid_creator)").
		Select("*, (uf_postid IS NOT NULL) AS favorited, ST_AsText(p_location) AS location").
		First(&post).Error
	if err != nil {
		return nil, httperr.New(err.Error(), http.StatusInternalServerError)
	}
	return &post, nil
}

func RemovePost(requesterID, postID string) httperr.Error {
	return ErrNotImplemented
}
