package db

import (
	"log"

	"github.com/alex-305/bookbackend/internal/db/helpers"
	"github.com/alex-305/bookbackend/internal/db/queries"
	"github.com/alex-305/bookbackend/internal/db/scan"
	"github.com/alex-305/bookbackend/internal/models"
)

func (db *DB) GetFollowerList(userusername, username string, o models.FollowSortOptions) ([]models.UserFollow, error) {
	q := queries.GetFollower()
	q += helpers.ListFormat(o)

	rows, err := db.Query(q, userusername, username)

	defer rows.Close()

	if err != nil {
		log.Printf("%s", err)
		return []models.UserFollow{}, err
	}
	return scan.Follows(rows)
}

func (db *DB) GetFollowingList(userusername, username string, o models.FollowSortOptions) ([]models.UserFollow, error) {
	q := queries.GetFollowing()
	q += helpers.ListFormat(o)

	rows, err := db.Query(q, userusername, username)

	defer rows.Close()

	if err != nil {
		log.Printf("%s", err)
		return []models.UserFollow{}, err
	}
	return scan.Follows(rows)
}
