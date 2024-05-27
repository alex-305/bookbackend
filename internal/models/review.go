package models

import "time"

type Review struct {
	ReviewID  ReviewID  `json:"reviewid" gorm:"primaryKey"`
	Content   string    `json:"content"`
	Rating    uint8     `json:"rating"`
	Username  string    `json:"username" gorm:"-"`
	UserID    UserID    `json:"-"`
	VolumeID  VolumeID  `json:"volumeid"`
	Post_date time.Time `json:"post_date" gorm:"autoCreateTime"`
	LikeCount uint      `json:"likecount"`
	IsLiked   bool      `json:"isliked" gorm:"-"`
}
type ReviewListStats struct {
	ReviewCount uint    `json:"reviewcount"`
	AvgRating   float32 `json:"avgrating"`
}

type ReviewStats struct {
	CommentCount uint `json:"commentcount"`
}

type ReviewID string
