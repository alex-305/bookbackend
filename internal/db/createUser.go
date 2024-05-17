package db

import (
	"github.com/alex-305/bookbackend/internal/models"
)

func (db *DB) CreateUser(creds models.Credentials) error {
	query := `INSERT INTO users(username, password, email) VALUES($1, $2, $3);`
	_, err := db.Exec(query, creds.Username, creds.Password, creds.Email)

	if err != nil {
		return err
	}
	return nil
}
