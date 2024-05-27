package db

import (
	"github.com/alex-305/bookbackend/internal/db/queries"
	"github.com/alex-305/bookbackend/internal/models"
)

func (db *DB) GetUserID(username models.Username) (models.UserID, error) {
	var userID string
	err := db.Table(queries.FromUser()).
		Select(queries.SelectUserID()).
		Where("username = ?", username).
		Scan(&userID).
		Error

	if err != nil {
		return "", err
	}

	return models.UserID(userID), nil
}

func (db *DB) GetUsername(userID models.UserID) (models.Username, error) {
	var username models.Username
	err := db.Table(queries.FromUser()).
		Select(queries.SelectUsername()).
		Where("userid = ?", userID).
		Scan(&username).
		Error

	if err != nil {
		return "", err
	}

	return username, nil
}
