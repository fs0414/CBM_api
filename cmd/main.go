package main

import (
  "github.com/soramar/CBM_api/router"
  "github.com/soramar/CBM_api/model/database"
)

func main() {
  database.DbInit()

  router := router.GetRouter()
  router.Run(":8080")
}
