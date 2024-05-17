package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"index;not null"`
	Email    string `gorm:"unique;index;not null"`
	Password string `gorm:"not null"`
}
