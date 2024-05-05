package usersvc

import (
	"github.com/alex-305/bookbackend/db"
	"github.com/alex-305/bookbackend/models"
)

func Get(username string, db *db.DB) (models.User, error) {
	return db.GetUser(username)
}
