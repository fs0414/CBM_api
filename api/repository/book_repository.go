package repository

import(
	"gorm.io/gorm"
	"github.com/soramar/CBM_api/model/schema"
	"github.com/soramar/CBM_api/model/database"
)

func getAll() *gorm.DB {
	var books schema.Book
	db := database.Db
	result := db.Find(&books)
	return result
}
