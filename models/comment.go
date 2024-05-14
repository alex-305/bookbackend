package models

type Comment struct {
	Content   string `json:"content"`
	Username  string `json:"username"`
	CommentID string `json:"commentid"`
	ReviewID  string `json:"reviewid"`
}
