package controller

import (
	"net/http"
	"github.com/soramar/CBM_api/api/repository"
	"github.com/gin-gonic/gin"
	"github.com/soramar/CBM_api/model/schema"
	"golang.org/x/crypto/bcrypt"
)

type UserResponse struct {
	ID    uint    `json:"id"`
	Name  string `json:"name"`
}

func GetUsers(context *gin.Context) {
	users, err := repository.GetAllUsers()

	if err != nil {
		context.JSON(200, nil)
		return
	}
	var responseUsers []UserResponse
	for _, user := range users {
			responseUser := UserResponse{
					ID:    user.ID,
					Name:  user.Name,
			}
			responseUsers = append(responseUsers, responseUser)
	}
	context.JSON(200, responseUsers)
}

func Register(c *gin.Context) {
	var user schema.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user.Password = hashedPassword

	err = repository.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}
