package db

import (
	"github.com/alex-305/bookbackend/models"
)

func (db *DB) GetUser(username string) (models.User, error) {
	row := db.QueryRow("SELECT * FROM users WHERE username = $1", username)

	var user models.User
	err := row.Scan(&user.Username, &user.Password, &user.Email, &user.Description, &user.Join_date)

	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
