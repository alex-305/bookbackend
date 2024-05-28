package models

import (
	"time"
)

type User struct {
	UserID         UserID    `json:"-" gorm:"primaryKey;column:userid"`
	Username       string    `json:"username"`
	Password       string    `json:"-"`
	Email          string    `json:"email" gorm:"unique;not null"`
	Description    string    `json:"description"`
	Join_date      time.Time `json:"join_date" gorm:"autoCreateTime"`
	FollowerCount  uint      `json:"followercount" gorm:"column:followercount"`
	FollowingCount uint      `json:"followingcount" gorm:"column:followingcount"`
	IsFollowing    bool      `json:"isfollowing" gorm:"column:isfollowing"`
}

func CredsToUser(creds Credentials) User {
	return User{
		Username: string(creds.Username),
		Password: creds.Password,
		Email:    creds.Email,
	}
}

type UserID string

type Username string
