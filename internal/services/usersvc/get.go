package usersvc

import (
	"github.com/alex-305/bookbackend/internal/auth/token"
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
)

func Get(username string, tok models.Token, db *db.DB) (models.User, error) {
	userusername, _ := token.Validate(tok)
	return db.GetUser(userusername, username)
}
