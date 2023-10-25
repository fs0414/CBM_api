package model

import (
	"time"
  "github.com/joho/godotenv"
  "gorm.io/gorm"
  "gorm.io/driver/mysql"
  "os"
)

type User struct {
	gorm.Model

	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Password  string `gorm:"not null"`
	Email     string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Book struct {
	gorm.Model

	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	ImageUrl  string `gorm:"not null"`
	Loanable  bool   `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type BorrowedBook struct {
	gorm.Model

	ID            uint      `gorm:"primaryKey"`
	UserID        uint      `gorm:"not null"`
	BookID        uint      `gorm:"not null"`
	CheckoutDate  time.Time `gorm:"not null"`
	ReturnDueDate time.Time `gorm:"not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type BorrowingWishList struct {
  gorm.Model

	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	BookID    uint      `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func DbInit() {
  godotenv.Load(".env")
  MYSQL_USER := os.Getenv("MYSQL_USER")
  MYSQL_PASSWORD := os.Getenv("MYSQL_PASSWORD")
  MYSQL_DATABASE := os.Getenv("MYSQL_DATABASE")
  MYSQL_IP := os.Getenv("MYSQL_IP")
  dsn := MYSQL_USER + ":" + MYSQL_PASSWORD + "@" + MYSQL_IP + "/" + MYSQL_DATABASE


  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil {
      panic("failed to connect to database")
  }
    
  db.AutoMigrate(&User{}, &Book{}, &BorrowedBook{}, &BorrowingWishList{})
}