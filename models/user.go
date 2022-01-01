package models

import "time"

type User struct {
	Id        uint
	Name      string `json:"name" form:"name"`
	Email     string `json:"email" form:"email" gorm:"unique"`
	Password  string `json:"password" form:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
