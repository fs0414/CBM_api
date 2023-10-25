package schema

import (
	"gorm.io/gorm"
)

type Role string

const (
	UserRole Role = "USER"
	AdminRole Role = "ADMIN"
)

type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255);uniqueIndex;not null"`
	Password string `gorm:"type:varchar(255);not null"`
	Role Role `gorm:"type:ENUM('USER', 'ADMIN');default:'USER';not null"`
	Books []Book
}

type Book struct {
	gorm.Model
	Title  string `gorm:"type:varchar(255);not null"`
	UserID uint
}