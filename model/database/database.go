package database

import (
	  "github.com/soramar/CBM_api/model/schema"
  	"github.com/joho/godotenv"
  	"gorm.io/gorm"
  	"gorm.io/driver/mysql"
  	"os"
    "fmt"
)

var Db *gorm.DB
var err error

func DbInit() {
  godotenv.Load(".env")
  MYSQL_USER := os.Getenv("MYSQL_USER")
  MYSQL_PASSWORD := os.Getenv("MYSQL_PASSWORD")
  MYSQL_DATABASE := os.Getenv("MYSQL_DATABASE")
  MYSQL_IP := os.Getenv("MYSQL_IP")
  dsn := MYSQL_USER + ":" + MYSQL_PASSWORD + "@" + MYSQL_IP + "/" + MYSQL_DATABASE

  Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil {
      fmt.Println("database connection faild", err)
      panic("failed to connect to database")
  }

  Db.AutoMigrate(&schema.User{}, &schema.Book{})
  fmt.Println("gorm db connect")

}

// func AutoMigration() {
//   fmt.Println("go auto migration")
//   Db.AutoMigrate(schema.User{}, schema.Book{})
// }