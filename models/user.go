package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"not null"`
	Email    string `json:"unique;not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"default:admin"`
}
