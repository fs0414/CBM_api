package repository

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/soramar/CBM_api/model/database"
	"github.com/soramar/CBM_api/model/schema"
	"gorm.io/gorm"
	"os"
	"time"
)

var jwtSecretKey = []byte(os.Getenv("ACCESS_SECRET_KEY"))

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
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

func IsTokenInvalid(tokenString string) (bool) {
	var invalidatedToken schema.InvalidatedToken
	err := database.Db.Where("token = ?", tokenString).First(&invalidatedToken).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false
		}
		return false
	}
	return true
}

func CreateInvalidateToken(invalidatedToken *schema.InvalidatedToken) error {
	err := database.Db.Create(&invalidatedToken).Error
	if err != nil {
		return err
	}
	return nil
}
