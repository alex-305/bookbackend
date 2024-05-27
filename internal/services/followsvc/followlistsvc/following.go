package followlistsvc

import (
	"github.com/alex-305/bookbackend/internal/auth/token"
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
)

func Following(username models.Username, tok models.Token, o models.FollowSortOptions, db *db.DB) ([]models.UserFollow, error) {
	userusername, err := token.Validate(tok)
	var useruserID models.UserID

	if err == nil {
		useruserID, err = db.GetUserID(userusername)
	}

	followinglist, err := db.GetFollowingList(useruserID, username, o)

	if err != nil {
		return []models.UserFollow{}, err
	}

	return followinglist, nil
}
