package repository

import (
	"github.com/soramar/CBM_api/model/schema"
	"github.com/soramar/CBM_api/model/database"
)

func CreateUser(user *schema.User) error {
	err := database.Db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUserByEmail(email string) (*schema.User, error) {
	var user schema.User
	err := database.Db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
