package model

import (
  "github.com/joho/godotenv"
  "gorm.io/gorm"
  "gorm.io/driver/mysql"
  "os"
)

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
    
  db.AutoMigrate()
}