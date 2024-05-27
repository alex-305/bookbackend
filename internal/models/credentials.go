package models

type Credentials struct {
	Username Username `json:"username" gorm:"username"`
	Password string   `json:"password" gorm:"password"`
	Email    string   `json:"email" gorm:"email"`
}
