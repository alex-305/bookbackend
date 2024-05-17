package usersvc

import (
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
)

func Get(username string, db *db.DB) (models.User, error) {
	return db.GetUser(username)
}
