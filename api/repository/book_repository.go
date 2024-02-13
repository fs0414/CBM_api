package repository

import (
	"github.com/soramar/CBM_api/model/database"
	"github.com/soramar/CBM_api/model/schema"
	"fmt"
)

func GetAllBooks() ([]schema.Book, error) {
	var books []schema.Book // ユーザー情報を格納するためのスライス
	db := database.Db

	// Findメソッドを使用して全ユーザー情報を取得
	result := db.Find(&books)
	if result.Error != nil {
		fmt.Println("エラー発生:", result.Error)
		return nil, result.Error // エラーがあれば、エラーを返す
	}

	fmt.Printf("取得したユーザー数: %d\n", result.RowsAffected)
	return books, nil // エラーがなければ、ユーザー情報のスライスとnilを返す
}
