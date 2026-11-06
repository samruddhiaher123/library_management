package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model

	Title             string `gorm:"not null"`
	Author            string
	ISBN              string `gorm:"unique;not null"`
	Publisher         string
	TotalQuantity     int
	AvailableQuantity int

	CategoryID uint
	Category   Category `gorm:"foreignKey:CategoryID"`

	BorrowRecords []BorrowRecord `gorm:"foreignKey:BookID"`
}
