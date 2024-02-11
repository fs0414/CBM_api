package repository

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/soramar/CBM_api/model/database"
	"github.com/soramar/CBM_api/model/schema"
	"gorm.io/gorm"
)

var jwtSecretKey = []byte(os.Getenv("ACCESS_SECRET_KEY"))

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

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

func IsEmailUnregistered(email string) bool {
	var user schema.User
	err := database.Db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return true
		}
		return true
	}
	return false
}

func GenerateToken(email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecretKey)

	return tokenString, err
}
