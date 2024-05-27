package followlistsvc

import (
	"github.com/alex-305/bookbackend/internal/auth/token"
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
)

func Follower(username models.Username, tok models.Token, o models.FollowSortOptions, db *db.DB) ([]models.UserFollow, error) {
	userusername, _ := token.Validate(tok)
	useruserid, err := db.GetUserID(userusername)
	if err != nil {
		return []models.UserFollow{}, err
	}
	userid, err := db.GetUserID(username)
	if err != nil {
		return []models.UserFollow{}, err
	}
	followerList, err := db.GetFollowerList(useruserid, userid, o)

	if err != nil {
		return []models.UserFollow{}, err
	}

	return followerList, nil
}
