package followlistsvc

import (
	"github.com/alex-305/bookbackend/internal/auth/token"
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
)

func Following(username string, tok models.Token, o models.FollowSortOptions, db *db.DB) ([]models.UserFollow, error) {
	userusername, _ := token.Validate(tok)
	return db.GetFollowingList(userusername, username, o)
}
