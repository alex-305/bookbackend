package db

import (
	"errors"

	"github.com/alex-305/bookbackend/internal/db/queries"
	"github.com/alex-305/bookbackend/internal/db/scan"
	"github.com/alex-305/bookbackend/internal/models"
)

func (db *DB) GetUser(userusername, username string) (models.User, error) {
	query := queries.GetUser()
	row := db.QueryRow(query, userusername, username)

	user, err := scan.User(row)

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

func (db *DB) CreateUser(creds models.Credentials) error {
	query := `INSERT INTO users(username, password, email) VALUES($1, $2, $3);`
	_, err := db.Exec(query, creds.Username, creds.Password, creds.Email)

	if err != nil {
		return err
	}
	return nil
}
