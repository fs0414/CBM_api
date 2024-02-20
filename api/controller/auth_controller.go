package controller

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/soramar/CBM_api/api/repository"
	"github.com/soramar/CBM_api/model/schema"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
	"strings"
)

func Login(c *gin.Context) {
	var loginRequest schema.LoginRequest

	err := c.BindJSON(&loginRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	var validationErrors = make(map[string][]string)

	if err := validate.Struct(loginRequest); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var errorMessage string
			fieldName := err.Field()
			tag := err.Tag()

			switch fieldName {
			case "Email":
				if tag == "required" {
					errorMessage = "メールアドレスは必須です"
				}
			case "Password":
				if tag == "required" {
					errorMessage = "パスワードは必須です"
				}
			}

			if errorMessage != "" {
				fieldName = strings.ToLower(fieldName)
				validationErrors[fieldName] = append(validationErrors[fieldName], errorMessage)
			}
		}

		c.JSON(http.StatusBadRequest, gin.H{"validation_error": validationErrors})
		return
	}

	user, err := repository.GetUserByEmail(loginRequest.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	tokenString, err := generateToken(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": tokenString})
}

func Logout(c *gin.Context) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
	})

	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}

func generateToken(c *gin.Context, user *schema.User) (string, error){
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("ACCESS_SECRET_KEY")))

	return tokenString, err
}
