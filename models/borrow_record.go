package models

import (
	"time"

	"gorm.io/gorm"
)

type BorrowRecord struct {
	gorm.Model

	MemberID uint
	Member   Member `gorm:"foreignKey:MemberID"`

	BookID uint
	Book   Book `gorm:"foreignKey:BookID"`

	IssuedDate time.Time
	DueDate    time.Time
	ReturnDate *time.Time

	FineAmount float64
	Status     string `gorm:"default:borrowed"`
}
