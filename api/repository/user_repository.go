package repository

import (
	"github.com/soramar/CBM_api/model/schema"
	"github.com/soramar/CBM_api/model/database"
	"gorm.io/gorm"
	"fmt"
)

// var user []schema.User = schema.User
// var db database.Db

func GetAll() *gorm.DB {
	var user schema.User
	db := database.Db
	fmt.Println("userRepo", user)
	result := db.Find(&user)
	fmt.Println("userRepo result", result)
	return result
}
