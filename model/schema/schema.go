package schema

import (
	"gorm.io/gorm"
	"time"
)

type Role string

const (
	UserRole Role = "USER"
	AdminRole Role = "ADMIN"
)

type User struct {
	gorm.Model
	Name 								string 	`gorm:"type:varchar(255);not null"`
	Email    						string 	`gorm:"type:varchar(255);uniqueIndex;not null"`
	Password 						string 	`gorm:"type:varchar(255);not null"`
	Role 								Role 		`gorm:"type:ENUM('USER', 'ADMIN');default:'USER';not null"`
	Books 							[]Book
	BorrowedBooks 			[]BorrowedBook
	BorrowingWishLists 	[]BorrowingWishList
}

type Book struct {
	gorm.Model
	UserId		uint
	User      User
	Title  		string `gorm:"type:varchar(255);not null"`
	ImageUrl  string `gorm:"type:varchar(255);not null"`
	Loanable  bool   `gorm:"not null"`
}

type BorrowedBook struct {
	gorm.Model
	UserID        uint      `gorm:"not null"`
	BookID        uint      `gorm:"not null"`
	CheckoutDate  time.Time `gorm:"not null"`
	ReturnDueDate time.Time `gorm:"not null"`
}

type BorrowingWishList struct {
	gorm.Model
	UserID    uint      `gorm:"not null"`
	BookID    uint      `gorm:"not null"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}