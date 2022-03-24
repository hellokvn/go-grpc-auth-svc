package models

import "gorm.io/gorm"

type Auth struct {
	gorm.Model
	Id       int64  `json:"id" gorm:"primaryKey"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
