package transactions

// Returns true if the requester can view favorites.
func CanViewFavorites(requester, userID string) (bool, error) {
	return requester == userID, nil
}

func CanModifyFavorites(requester, userID string) (bool, error) {
	return requester == userID, nil
}

func AddFavorite(requester, userID, postID string) error {
	permission, err := CanModifyFavorites(requester, userID)
	if err != nil {
		return err
	} else if !permission {
		return ErrInsufficientPermission
	}
	return db.Exec("INSERT INTO user_favorites (uf_userid, uf_postid) VALUES (?, ?)", requester, postID).Error
}

func GetFavorites(requester, userID string) ([]string, error) {
	permission, err := CanViewFavorites(requester, userID)
	if err != nil {
		return nil, err
	} else if !permission {
		return nil, ErrInsufficientPermission
	}
	var favorites []struct {
		PostID string `gorm:"column:uf_postid"`
	}
	err = db.Table("user_favorites").Where("uf_userid = ?", userID).Find(&favorites).Error
	if err != nil {
		return nil, err
	}
	postIDs := make([]string, len(favorites))
	for i, favorite := range favorites {
		postIDs[i] = favorite.PostID
	}
	return postIDs, nil
}

func RemoveFavorite(requester, userID, postID string) error {
	permission, err := CanModifyFavorites(requester, userID)
	if err != nil {
		return err
	} else if !permission {
		return ErrInsufficientPermission
	}
	return db.Exec("DELETE FROM user_favorites WHERE uf_userid = ? AND uf_postid = ?", requester, postID).Error
}
