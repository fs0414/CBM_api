package controller

import (
	"net/http"
	"github.com/soramar/CBM_api/api/repository"
	"github.com/gin-gonic/gin"
	"github.com/soramar/CBM_api/model/schema"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

type UserResponse struct {
	ID    uint    `json:"id"`
	Name  string `json:"name"`
}

func GetUsers(context *gin.Context) {
	users, err := repository.GetAllUsers()

	if err != nil {
		context.JSON(http.StatusOK, nil)
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
	context.JSON(http.StatusOK, responseUsers)
}

func Register(c *gin.Context) {
	var user schema.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()

	if err := validate.Struct(user); err != nil {
		var errorMessages []string

		for _, err := range err.(validator.ValidationErrors) {
			var errorMessage string

			fieldName := err.Field()
			typ := err.Tag()
			param := err.Param()

			switch fieldName {
				case "Name":
					errorMessage = "名前は必須です"
				case "Email":
					switch typ {
					case "required":
						errorMessage = "メールアドレスは必須です"
					case "email":
						errorMessage = "メールアドレスのフォーマットで登録してください"
					}
				case "Password":
					switch typ {
					case "required":
						errorMessage = "パスワードは必須です"
					case "min":
						errorMessage = fmt.Sprintf("パスワードは%s文字以上で登録してください", param)
					}
				case "Role":
					errorMessage = "権限は必須です"
			}
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{"validation_error": errorMessages})
		return
	}

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user.Password = hashedPassword

	if err := repository.CreateUser(&user); err != nil {
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
