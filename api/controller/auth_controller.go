package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/soramar/CBM_api/api/repository"
	"github.com/soramar/CBM_api/model/schema"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var loginRequest schema.LoginRequest
	err := c.BindJSON(&loginRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	// パスワードのハッシュ値を比較
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// ログイン成功の処理
	// 例えば、トークンを生成して返すなど

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func Logout(c *gin.Context) {

}