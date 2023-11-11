package controller

import (
	"net/http"
	"github.com/soramar/CBM_api/api/repository"
	"github.com/gin-gonic/gin"
	"github.com/soramar/CBM_api/model/schema"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(context *gin.Context) {
	users := repository.GetAll() 
	context.JSON(200, users)
}

func Register(c *gin.Context) {
	var user schema.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// パスワードをハッシュ化
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// ハッシュ化されたパスワードをセット
	user.Password = hashedPassword

	err = repository.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func hashPassword(password string) (string, error) {
	// パスワードをbcryptでハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}