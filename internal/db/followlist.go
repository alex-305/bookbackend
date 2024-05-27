package db

import (
	"log"

	"github.com/alex-305/bookbackend/internal/models"
)

func (db *DB) GetFollowerList(userid models.UserID, followingUserID models.UserID, o models.FollowSortOptions) ([]models.UserFollow, error) {
	var followers []models.UserFollow
	err :=
		db.Table(`user_follows_user uf1`).
			Select(`uf1.followerid AS userid, CASE WHEN uf2.followerid IS NOT NULL THEN TRUE ELSE FALSE END AS followed`).
			Joins(`LEFT JOIN user_follows_user uf2 ON uf1.followedid=uf2.followerid AND uf2.followedid=?`, followingUserID).
			//Joins(queries.JoinForUsername("uf1")).
			Where(`uf1.followerid=?`, userid).
			Order(string(o.By) + " " + string(o.Direction)).
			Limit(int(o.Limit)).
			Offset(int(o.Limit * o.Page)).
			Scan(&followers).
			Error
	if err != nil {
		log.Printf("%s", err)
		return []models.UserFollow{}, err
	}

	return followers, nil
}

func (db *DB) GetFollowingList(username models.UserID, follower models.Username, o models.FollowSortOptions) ([]models.UserFollow, error) {
	var followings []models.UserFollow
	err :=
		db.Table("user_follows_user uf1").
			Select("uf1.followerid, CASE WHEN uf2.followerid IS NOT NULL THEN TRUE ELSE FALSE END AS followed").
			Joins("LEFT JOIN user_follows_user uf2 ON uf1.followedid=uf2.followerid AND uf2.followedid = ?", follower).
			Where("uf1.followerid = ?", username).
			Order(string(o.By) + " " + string(o.Direction)).
			Limit(int(o.Limit)).
			Offset(int(o.Limit * o.Page)).
			Scan(&followings).
			Error
	if err != nil {
		log.Printf("%s", err)
		return []models.UserFollow{}, err
	}

	return followings, nil
}
