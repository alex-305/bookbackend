package db

import (
	"errors"

	"github.com/alex-305/bookbackend/models"
)

func (db *DB) GetUser(username string) (models.User, error) {
	row := db.QueryRow("SELECT username, email, description, join_date FROM users WHERE username = $1", username)

	var user models.User
	err := row.Scan(&user.Username, &user.Email, &user.Description, &user.Join_date)

	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (db *DB) GetCredentials(username string) (models.Credentials, error) {
	row := db.QueryRow("SELECT username, email, password FROM users WHERE username = $1", username)
	var creds models.Credentials

	err := row.Scan(&creds.Username, &creds.Email, &creds.Password)

	if err != nil {
		return models.Credentials{}, errors.New("unable to get credentials")
	}

	return creds, nil
}
