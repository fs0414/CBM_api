package repository

import (
	"github.com/soramar/CBM_api/model/schema"
	"github.com/soramar/CBM_api/model/database"
	"fmt"
)

func GetAllUsers() ([]schema.User, error) {
	var users []schema.User
	db := database.Db

	result := db.Find(&users)
	if result.Error != nil {
		fmt.Println("エラー発生:", result.Error)
		return nil, result.Error
	}

	fmt.Printf("取得したユーザー数: %d\n", result.RowsAffected)
	return users, nil
}

func CreateUser(user *schema.User) error {
	err := database.Db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}
