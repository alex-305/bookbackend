package reviewlistsvc

import (
	"github.com/alex-305/bookbackend/internal/auth/token"
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
)

func GetUser(username models.Username, tok models.Token, o models.ReviewSortOptions, db *db.DB) ([]models.Review, error) {
	var useruserid models.UserID
	userusername, err := token.Validate(tok)

	if err == nil {
		useruserid, _ = db.GetUserID(userusername)
	}

	userid, err := db.GetUserID(username)

	if err != nil {
		return []models.Review{}, err
	}

	rl, err := db.GetUserReviewList(useruserid, userid, o)

	if err != nil {
		return []models.Review{}, err
	}

	return rl, nil

}
