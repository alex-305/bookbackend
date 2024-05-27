package tables

import (
	"time"

	"github.com/alex-305/bookbackend/internal/models"
)

type User_follows_user struct {
	FollowerID  models.UserID `gorm:"primaryKey"`
	Follow_date time.Time     `gorm:"autoCreateTime"`
	FollowedID  models.UserID `gorm:"primaryKey"`
}
