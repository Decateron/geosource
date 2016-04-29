package transactions

import (
	"net/http"

	"github.com/joshheinrichs/httperr"
)

// CanViewFavorites returns true if the requester has permission to view the
// favorites of the given user, false otherwise. Returns an error if a result
// cannot be determined.
func CanViewFavorites(requesterID, userID string) (bool, httperr.Error) {
	return requesterID == userID, nil
}

// CanModifyFavorites returns true if the requester has permission to modify the
// favorites of the given user, false otherwise. Returns an error if a result
// cannot be determined.
func CanModifyFavorites(requesterID, userID string) (bool, httperr.Error) {
	return requesterID == userID, nil
}

// AddFavorite adds the given post to the set of favorites for the user with ID
// userID. This transaction is executed under the permission level of the
// requester. Returns an error if the requester does not have sufficient
// permission, or if some other error occurs within the database.
func AddFavorite(requesterID, userID, postID string) httperr.Error {
	permission, httpErr := CanModifyFavorites(requesterID, userID)
	if httpErr != nil {
		return httpErr
	} else if !permission {
		return ErrInsufficientPermission
	}
	err := db.Exec("INSERT INTO user_favorites (uf_userid, uf_postid) VALUES (?, ?)", requesterID, postID).Error
	if err != nil {
		return httperr.New(err.Error(), http.StatusInternalServerError)
	}
	return nil
}

// GetFavorites returns the set of favorites for the user with ID userID. This
// transaction is executed under the permission level of the requester. Returns
// an error if the requester does not have sufficient permission, or if some
// other error occurs within the database.
func GetFavorites(requesterID, userID string) ([]string, httperr.Error) {
	permission, httpErr := CanViewFavorites(requesterID, userID)
	if httpErr != nil {
		return nil, httpErr
	} else if !permission {
		return nil, ErrInsufficientPermission
	}
	var favorites []struct {
		PostID string `gorm:"column:uf_postid"`
	}
	err := db.Table("user_favorites").Where("uf_userid = ?", userID).Find(&favorites).Error
	if err != nil {
		return nil, httperr.New(err.Error(), http.StatusInternalServerError)
	}
	postIDs := make([]string, len(favorites))
	for i, favorite := range favorites {
		postIDs[i] = favorite.PostID
	}
	return postIDs, nil
}

// RemoveFavorite removes the given post from the user with ID userID's set of
// favorites. This transaction is executed under the permission level of the
// requester. Returns an error if the requester does not have sufficient
// permission, or if some other error occurs within the database.
func RemoveFavorite(requesterID, userID, postID string) httperr.Error {
	permission, httpErr := CanModifyFavorites(requesterID, userID)
	if httpErr != nil {
		return httpErr
	} else if !permission {
		return ErrInsufficientPermission
	}
	err := db.Exec("DELETE FROM user_favorites WHERE uf_userid = ? AND uf_postid = ?", requesterID, postID).Error
	if err != nil {
		return httperr.New(err.Error(), http.StatusInternalServerError)
	}
	return nil
}
