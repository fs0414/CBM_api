package main

import (
  "fmt"
  "net/http"
  "os"
  "github.com/joho/godotenv"
  "gorm.io/gorm"
  "gorm.io/driver/mysql"
)

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
    
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, MYSQL_USER + "\n")
        fmt.Fprintf(w, dsn + "\n")
        // fmt.Fprintf(w, db + "\n")
        fmt.Fprintf(w, "hello cbmApi to golang air")
      })

  http.ListenAndServe(":8080", nil)

}
