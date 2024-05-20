package reviewlistsvc

import (
	"github.com/alex-305/bookbackend/internal/auth/token"
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
)

func GetUser(username string, tok models.Token, o models.ReviewSortOptions, db *db.DB) ([]models.Review, error) {
	userusername, _ := token.Validate(tok)
	return db.GetUserReviewList(userusername, username, o)
}
