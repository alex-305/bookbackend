package models

import "time"

type User struct {
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	Description    string    `json:"description"`
	Join_date      time.Time `json:"join_date"`
	FollowerCount  uint      `json:"followercount"`
	FollowingCount uint      `json:"followingcount"`
	IsFollowing    bool      `json:"isfollowing"`
}
