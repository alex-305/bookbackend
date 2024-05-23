package models

type UserFollow struct {
	Username    string `json:"username"`
	IsFollowing bool   `json:"isfollowing"`
}
