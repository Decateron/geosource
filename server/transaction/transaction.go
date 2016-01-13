package transaction

type User struct{}
type Channel struct{}
type Post struct{}
type Comment struct{}

// Adds the given User to the database, or returns an error if the insertion
// was unsuccessful.
func AddUser(requester, user *User) error {
	return nil
}

func GetUsername(email string) (*string, error) {
	return nil, nil
}

// Returns a user with the given username or nil if they do not exist. An error
// is returned if the database was accessed unsuccessfully.
func GetUser(requester, username string) (*User, error) {
	return nil, nil
}

func RemoveUser(requester, username string) error {
	return nil
}

func AddAdmin(requester, username string) error {
	return nil
}

func IsAdmin(requester, username string) (bool, error) {
	return false, nil
}

func RemoveAdmin(requester, username string) error {
	return nil
}

func AddChannel(requester, channel *Channel) error {
	return nil
}

func GetChannel(requester, channelname string) (*Channel, error) {
	return nil, nil
}

func RemoveChannel(requester, channelname string) error {
	return nil
}

func IsChannelCreator(requester, username, channelname string) (bool, error) {
	return false, nil
}

func AddModerator(requester, username, channelname string) error {
	return nil
}

func GetModerators(requester, channelname string) ([]string, error) {
	return nil, nil
}

func IsModerator(requester, username, channelname string) (bool, error) {
	return false, nil
}

func RemoveModerator(requester, username, channelname string) error {
	return nil
}

func AddViewer(requester, username, channelname string) error {
	return nil
}

func GetViewers(requester, channelname string) ([]string, error) {
	return nil, nil
}

func IsViewer(requester, username, channelname string) (bool, error) {
	return false, nil
}

func RemoveViewer(requester, username, channelname string) error {
	return nil
}

func AddBan(requester, username, channelname string) error {
	return nil
}

func GetBans(requester, channelname string) ([]string, error) {
	return nil, nil
}

func IsBanned(requester, username, channelname string) (bool, error) {
	return false, nil
}

func RemoveBan(requester, username, channelname string) error {
	return nil
}

func AddPost(requester, post *Post) error {
	return nil
}

// func GetPosts

func RemovePost(requester, postId string) error {
	return nil
}

func IsPostCreator(requester, username, postId string) (bool, error) {
	return false, nil
}

func AddComment(requester, comment *Comment) error {
	return nil
}

// func GetComments

func RemoveComment(requester, commentId string) error {
	return nil
}

func IsCommentCreator(requester, username, commentId string) (bool, error) {
	return false, nil
}

func AddFavorite(requester, username, postId string) error {
	return nil
}

// func GetFavorites

func RemoveFavorite(requester, username, postId string) error {
	return nil
}

func AddSubscription(requester, username, channelname string) error {
	return nil
}

// func GetSubscriptions

func RemoveSubscription(requester, username, channelname string) error {
	return nil
}
