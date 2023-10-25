package main

import (
  "github.com/soramar/CBM_api/router"
  // "github.com/soramar/CBM_api/api/controller"
  "github.com/soramar/CBM_api/model/database"
  // "github.com/soramar/CBM_api/model/schema"
  // "gorm.io/gorm"
  // "fmt"
)

func main() {
  database.DbInit()

  router := router.GetRouter()
  router.Run(":8080")
}