package transaction

import (
	"../config"
	"../types"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var db gorm.DB

func Init(config *config.Config) (err error) {
	db, err = gorm.Open("postgres", fmt.Sprintf("host=%s dbname=%s user=%s password=%s",
		config.Database.Host, config.Database.Database, config.Database.User, config.Database.Password))
	if err != nil {
		return err
	}
	return nil
}

// Adds the given User to the database, or returns an error if the insertion
// was unsuccessful.
func AddUser(user *types.User) error {
	err := db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUserByEmail(email string) (*types.User, error) {
	var user types.User
	err := db.Where("u_email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.RecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func GetUserById(id string) (*types.User, error) {
	var user types.User
	err := db.Where("u_userid = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUsername(email string) (*string, error) {
	return nil, nil
}

// Returns a user with the given username or nil if they do not exist. An error
// is returned if the database was accessed unsuccessfully.
func GetUser(requesterUid, uid string) (*types.User, error) {
	return nil, nil
}

func RemoveUser(requesterUid, uid string) error {
	return nil
}

func AddAdmin(requesterUid, uid string) error {
	return nil
}

func IsAdmin(requesterUid, uid string) (bool, error) {
	return false, nil
}

func RemoveAdmin(requesterUid, uid string) error {
	return nil
}

func AddChannel(requesterUid, channel *types.Channel) error {
	return nil
}

func GetChannel(requesterUid, channelname string) (*types.Channel, error) {
	return nil, nil
}

func RemoveChannel(requesterUid, channelname string) error {
	return nil
}

func IsChannelCreator(requesterUid, uid, channelname string) (bool, error) {
	return false, nil
}

func AddModerator(requesterUid, uid, channelname string) error {
	return nil
}

func GetModerators(requesterUid, channelname string) ([]string, error) {
	return nil, nil
}

func IsModerator(requesterUid, uid, channelname string) (bool, error) {
	return false, nil
}

func RemoveModerator(requesterUid, uid, channelname string) error {
	return nil
}

func AddViewer(requesterUid, uid, channelname string) error {
	return nil
}

func GetViewers(requesterUid, channelname string) ([]string, error) {
	return nil, nil
}

func IsViewer(requesterUid, uid, channelname string) (bool, error) {
	return false, nil
}

func RemoveViewer(requesterUid, uid, channelname string) error {
	return nil
}

func AddBan(requesterUid, uid, channelname string) error {
	return nil
}

func GetBans(requesterUid, channelname string) ([]string, error) {
	return nil, nil
}

func IsBanned(requesterUid, uid, channelname string) (bool, error) {
	return false, nil
}

func RemoveBan(requesterUid, uid, channelname string) error {
	return nil
}

func AddPost(requesterUid, post *types.Post) error {
	return nil
}

// func GetPosts

func RemovePost(requesterUid, postId string) error {
	return nil
}

func IsPostCreator(requesterUid, uid, postId string) (bool, error) {
	return false, nil
}

func AddComment(requesterUid, comment *types.Comment) error {
	return nil
}

// func GetComments

func RemoveComment(requesterUid, commentId string) error {
	return nil
}

func IsCommentCreator(requesterUid, uid, commentId string) (bool, error) {
	return false, nil
}

func AddFavorite(requesterUid, uid, postId string) error {
	return nil
}

// func GetFavorites

func RemoveFavorite(requesterUid, uid, postId string) error {
	return nil
}

func AddSubscription(requesterUid, uid, channelname string) error {
	return nil
}

// func GetSubscriptions

func RemoveSubscription(requesterUid, uid, channelname string) error {
	return nil
}
