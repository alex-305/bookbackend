package reviewlistsvc

import (
	"github.com/alex-305/bookbackend/internal/auth/token"
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
)

func GetBook(volumeid models.VolumeID, tok models.Token, o models.ReviewSortOptions, db *db.DB) ([]models.Review, error) {
	username, _ := token.Validate(tok)
	userid, _ := db.GetUserID(username)

	reviews, err := db.GetBookReviewList(userid, volumeid, o)

	if err != nil {
		return []models.Review{}, err
	}
	return reviews, nil
}
