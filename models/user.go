package models

import "time"

type User struct {
	Username    string
	Password    string
	Email       string
	Description string
	Join_date   time.Time
}
