package reviewlistsvc

import (
	"github.com/alex-305/bookbackend/internal/auth/token"
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
)

func GetBook(volumeid string, tok models.Token, o models.ReviewSortOptions, db *db.DB) ([]models.Review, error) {
	username, _ := token.Validate(tok)
	return db.GetBookReviewList(username, volumeid, o)
}
