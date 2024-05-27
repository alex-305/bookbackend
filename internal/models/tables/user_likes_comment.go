package tables

import (
	"time"

	"github.com/alex-305/bookbackend/internal/models"
)

type User_likes_comment struct {
	UserID    models.UserID    `gorm:"primaryKey"`
	Like_date time.Time        `gorm:"autoCreateTime"`
	CommentID models.CommentID `gorm:"primaryKey"`
}
