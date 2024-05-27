package db

import (
	"errors"

	"github.com/alex-305/bookbackend/internal/db/queries"
	"github.com/alex-305/bookbackend/internal/models"
)

func (db *DB) GetUser(useruserID models.UserID, userID models.UserID) (models.User, error) {
	var user models.User
	err := db.Table(queries.FromUser()).
		Select(queries.SelectUser()).
		Joins(queries.JoinUserFollows(), useruserID).
		Where("userid = ?", userID).
		Scan(&user).
		Error

	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (db *DB) GetCredentials(userid models.UserID) (models.Credentials, error) {
	var creds models.Credentials
	err := db.Table(queries.FromUser()).
		Select(queries.SelectCredentials()).
		Where("userid = ?", userid).
		Scan(&creds).
		Error

	if err != nil {
		return models.Credentials{}, errors.New("unable to get credentials")
	}

	return creds, nil
}

func (db *DB) CreateUser(creds models.Credentials) error {

	err := db.Table("users").Create(&creds).Error

	if err != nil {
		return err
	}
	return nil
}
