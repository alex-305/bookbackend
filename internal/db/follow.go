package db

import (
	"errors"

	"github.com/alex-305/bookbackend/internal/db/queries"
	"github.com/alex-305/bookbackend/internal/models"
	"github.com/alex-305/bookbackend/internal/models/tables"
)

func (db *DB) PostFollow(followerID models.UserID, followedID models.UserID) error {
	tx := db.Begin()

	newFollow := tables.User_follows_user{
		FollowerID: followerID,
		FollowedID: followedID,
	}

	result := tx.Create(&newFollow)

	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	result = queries.ChangeFollowerCount(followedID, tx, 1)

	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	result = queries.ChangeFollowingCount(followerID, tx, 1)

	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	tx.Commit()

	return nil

}

func (db *DB) DeleteFollow(followerID models.UserID, followedID models.UserID) error {
	tx := db.Begin()
	result := tx.Delete(&tables.User_follows_user{}, "followerID=? AND followedID=?", followerID, followedID)

	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	if result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("could not delete from user_follows_user")
	}

	result = queries.ChangeFollowerCount(followedID, tx, -1)

	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	result = queries.ChangeFollowingCount(followerID, tx, -1)

	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	tx.Commit()
	return nil
}
