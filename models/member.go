package models

import "gorm.io/gorm"

type Member struct {
	gorm.Model

	Name    string `gorm:"not null"`
	Email   string `gorm:"unique;not null"`
	Phone   string
	Address string

	BorrowRecords []BorrowRecord `gorm:"foreignKey:MemberID"`
}
