package tables

import (
	"time"

	"github.com/alex-305/bookbackend/internal/models"
)

type User_likes_review struct {
	UserID    models.UserID   `gorm:"primaryKey"`
	Like_date time.Time       `gorm:"autoCreateTime"`
	ReviewID  models.ReviewID `gorm:"primaryKey"`
}
