package usersvc

import (
	"github.com/alex-305/bookbackend/internal/auth/token"
	"github.com/alex-305/bookbackend/internal/db"
	"github.com/alex-305/bookbackend/internal/models"
)

func Get(username models.Username, tok models.Token, db *db.DB) (models.User, error) {
	userusername, err := token.Validate(tok)

	var useruserid models.UserID
	if err == nil {
		useruserid, _ = db.GetUserID(userusername)
	}

	userid, err := db.GetUserID(username)

	user, err := db.GetUser(useruserid, userid)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
