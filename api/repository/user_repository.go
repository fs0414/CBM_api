package repository

import (
	"github.com/soramar/CBM_api/model/schema"
	"github.com/soramar/CBM_api/model/database"
	"fmt"
)

// var user []schema.User = schema.User
// var db database.Db

func GetAllUsers() ([]schema.User, error) {
	var users []schema.User // ユーザー情報を格納するためのスライス
	db := database.Db

	// Findメソッドを使用して全ユーザー情報を取得
	result := db.Find(&users)
	if result.Error != nil {
		fmt.Println("エラー発生:", result.Error)
		return nil, result.Error // エラーがあれば、エラーを返す
	}

	fmt.Printf("取得したユーザー数: %d\n", result.RowsAffected)
	return users, nil // エラーがなければ、ユーザー情報のスライスとnilを返す
}
