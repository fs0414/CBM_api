package main

import (
  "fmt"
  "net/http"
  "os"
  "github.com/joho/godotenv"
  "gorm.io/gorm"
  "gorm.io/driver/mysql"
  "github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
  r := gin.Default()
  r.GET("/", func(c *gin.Context) {
    c.String(http.StatusOK, "hello cbm to golang")
  })

  return r
}

func main() {
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
    
  fmt.Println("database connect success")

  db.AutoMigrate()

  router := GetRouter()
  router.Run(":8080")
}