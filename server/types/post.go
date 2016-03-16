package types

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/joshheinrichs/geosource/server/types/fields"
)

const (
	MAX_POSTTITLE_LEN = 140
)

// General meta information about a post.
type PostInfo struct {
	ID        string    `json:"id" gorm:"column:p_postid"`
	CreatorID string    `json:"creator" gorm:"column:p_userid_creator"`
	Channel   string    `json:"channel" gorm:"column:p_channelname"`
	Title     string    `json:"title" gorm:"column:p_title"`
	Thumbnail string    `json:"thumbnail" gorm:"column:p_thumbnail"`
	Time      time.Time `json:"time" gorm:"column:p_time"`
	Location  Location  `json:"location" gorm:"column:p_location" sql:"type:POINT NOT NULL"`
}

// Returns the name of the postInfo's corresponding table in the database.
func (postInfo *PostInfo) TableName() string {
	return "posts"
}

// Validates the post info. Returns an error if any fields are invalid, nil
// otherwise.
func (postInfo *PostInfo) Validate() error {
	postInfo.Title = strings.TrimSpace(postInfo.Title)
	if len(postInfo.Title) == 0 {
		return errors.New("Post title cannot be empty.")
	} else if len(postInfo.Title) > MAX_POSTTITLE_LEN {
		return errors.New(fmt.Sprintf("Length of post title cannot exceed %i characters", MAX_POSTTITLE_LEN))
	}
	err := postInfo.Location.Validate()
	if err != nil {
		return err
	}
	return nil
}

type Post struct {
	PostInfo
	Fields fields.Fields `json:"fields" gorm:"column:p_fields" sql:"type:JSONB NOT NULL"`
}

// Validates the post. Returns an error if any fields are invalid, nil
// otherwise.
func (post *Post) Validate() error {
	err := post.PostInfo.Validate()
	if err != nil {
		return err
	}
	err = post.Fields.ValidateValues()
	if err != nil {
		return err
	}
	return nil
}

// Generates a thumbnail for the post, attempting to use an image within the
// post if one exists. This function assumes that the images from the post have
// been validated and written to storage.
func (post *Post) GenerateThumbnail() error {
	// TODO: Assign default thumbnail
	for _, field := range post.Fields {
		imagesValue, ok := field.Value.(*fields.ImagesValue)
		if ok && imagesValue.IsComplete() {
			thumbnail, err := imagesValue.GenerateThumbnail()
			if err != nil {
				return err
			}
			post.Thumbnail = thumbnail
			return nil
		}
	}
	post.Thumbnail = fields.MediaDir + fields.ThumbnailDir + "default.svg"
	return nil
}

// Returns the name of the post's corresponding table in the database.
func (post *Post) TableName() string {
	return "posts"
}
