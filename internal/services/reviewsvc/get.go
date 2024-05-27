package reviewsvc

import (
	"github.com/alex-305/bookbackend/internal/auth/token"
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
)

func Get(reviewid models.ReviewID, tok models.Token, db *db.DB) (models.Review, error) {
	username, _ := token.Validate(tok)
	userID, _ := db.GetUserID(username)

	return db.GetReview(userID, reviewid)
}
