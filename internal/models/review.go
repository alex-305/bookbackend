package models

import "time"

type Review struct {
	ReviewID  ReviewID  `json:"reviewid" gorm:"primaryKey;column:reviewid"`
	Content   string    `json:"content"`
	Rating    uint8     `json:"rating"`
	Username  string    `json:"username" gorm:"-"`
	UserID    UserID    `json:"-" gorm:"column:userid"`
	VolumeID  VolumeID  `json:"volumeid" gorm:"column:volumeid"`
	Post_date time.Time `json:"post_date" gorm:"autoCreateTime"`
	LikeCount uint      `json:"likecount" gorm:"column:likecount"`
	IsLiked   bool      `json:"isliked" gorm:"column:isliked"`
}

type NewReview struct {
	ReviewID ReviewID `json:"reviewid" gorm:"primaryKey;column:reviewid"`
	Content  string   `json:"content"`
	Rating   uint8    `json:"rating"`
	Username string   `json:"username" gorm:"-"`
	UserID   UserID   `json:"-" gorm:"column:userid"`
	VolumeID VolumeID `json:"volumeid" gorm:"column:volumeid"`
}
type ReviewListStats struct {
	ReviewCount uint    `json:"reviewcount"`
	AvgRating   float32 `json:"avgrating"`
}

type ReviewStats struct {
	CommentCount uint `json:"commentcount"`
}

type ReviewID string
