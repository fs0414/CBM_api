package repository

import (
	"github.com/soramar/CBM_api/model/schema"
	"github.com/soramar/CBM_api/model/database"
	"time"
	"github.com/dgrijalva/jwt-go"
	"os"
)

var jwtSecretKey = []byte(os.Getenv("ACCESS_SECRET_KEY"))

type Claims struct {
	Email string `json:"Email"`
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