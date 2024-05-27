package db

import (
	"log"

	"github.com/alex-305/bookbackend/internal/db/queries"
	"github.com/alex-305/bookbackend/internal/models"
)

func (db DB) GetUserReviewList(useruserID models.UserID, userID models.UserID, o models.ReviewSortOptions) ([]models.Review, error) {
	var rev []models.Review
	err := db.Table(queries.FromReview()).
		Select(queries.SelectReview()).
		Joins(queries.JoinReviewLikes(), useruserID).
		//Joins(queries.JoinForUsername(queries.ReviewTableName())).
		Where(`r.userid=?`, userID).
		Order(string(o.By) + " " + string(o.Direction)).
		Limit(int(o.Limit)).
		Offset(int(o.Limit * o.Page)).
		Scan(&rev).
		Error

	if err != nil {
		log.Printf("%s", err)
		return []models.Review{}, err
	}
	return rev, nil
}

func (db DB) GetBookReviewList(userID models.UserID, volumeID models.VolumeID, o models.ReviewSortOptions) ([]models.Review, error) {
	var rev []models.Review

	err := db.Table(queries.FromReview()).
		Select(queries.SelectReview()).
		Joins(queries.JoinReviewLikes(), userID).
		//Joins(queries.JoinForUsername(queries.ReviewTableName())).
		Where(`r.volumeid=?`, volumeID).
		Order(string(o.By) + " " + string(o.Direction)).
		Limit(int(o.Limit)).
		Offset(int(o.Limit * o.Page)).
		Scan(&rev).
		Error

	if err != nil {
		log.Printf("%s", err)
		return []models.Review{}, err
	}
	return rev, nil
}

func (db DB) GetFollowingReviewList(userID models.UserID, o models.ReviewSortOptions) ([]models.Review, error) {
	var rev []models.Review
	err := db.Table(queries.FromReview()).
		Select(queries.SelectReview()).
		Joins(queries.JoinReviewFollowing()).
		Joins(queries.JoinReviewLikes(), userID).
		//Joins(queries.JoinForUsername(queries.ReviewTableName())).
		Where(`ufu.followerid = ?`, userID).
		Order(string(o.By) + " " + string(o.Direction)).
		Limit(int(o.Limit)).
		Offset(int(o.Limit * o.Page)).
		Scan(&rev).
		Error

	if err != nil {
		log.Printf("%s", err)
		return []models.Review{}, err
	}

	return rev, nil
}
