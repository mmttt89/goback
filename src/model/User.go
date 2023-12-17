package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string
	Email     string `gorm:"unique"`
	User_Role Role   `gorm:"default:user"`
}
