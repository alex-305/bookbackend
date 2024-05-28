package models

import "time"

type Comment struct {
	CommentID CommentID `json:"commentid" gorm:"primaryKey"`
	Content   string    `json:"content"`
	UserID    UserID    `json:"-"`
	Username  Username  `json:"username" gorm:"-"`
	ReviewID  ReviewID  `json:"reviewid"`
	Post_date time.Time `json:"post_date" gorm:"autoCreateTime"`
	LikeCount uint      `json:"likecount"`
	IsLiked   bool      `json:"isliked" gorm:"-"`
}

type NewComment struct {
	CommentID CommentID `json:"commentid" gorm:"primaryKey"`
	Content   string    `json:"content"`
	UserID    UserID    `json:"-"`
	Username  Username  `json:"username" gorm:"-"`
	ReviewID  ReviewID  `json:"reviewid"`
}

type CommentStats struct {
	CommentCount uint `json:"commentcount"`
}

type CommentID string
