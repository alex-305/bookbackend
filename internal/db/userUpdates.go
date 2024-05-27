package db

import "github.com/alex-305/bookbackend/internal/models"

func (db DB) UpdateUserDescription(userID models.UserID, description string) error {
	err := db.Model(&models.User{}).Where("userID = ?", userID).Update("description", description).Error
	if err != nil {
		return err
	}
	return err
}

func (db DB) UpdateUserPassword(userID models.UserID, password string) error {
	err := db.Model(&models.User{}).Where("userID = ?", userID).Update("password", password).Error
	if err != nil {
		return err
	}
	return err
}
