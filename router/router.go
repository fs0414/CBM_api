package router

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/soramar/CBM_api/api/controller"
	"github.com/soramar/CBM_api/api/repository"
	"net/http"
	"os"
)

func UserRegisteredMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("jwt")

		if err != nil || tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("ACCESS_SECRET_KEY")), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			email, ok := claims["email"].(string)

			if !ok {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
				return
			}

			IsEmailUnregistered := repository.IsEmailUnregistered(email)

			if IsEmailUnregistered {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not registered"})
				return
			}

		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token or user not registered"})
		}
	}
}

func GetRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", controller.SayHello)
	api := r.Group("/api", UserRegisteredMiddleware())
	auth := api.Group("auth")
	{
		auth.POST("/register", controller.Register)
		r.POST("api/login", controller.Login)
		r.POST("api/logout", controller.Logout)
		api.GET("/users", controller.GetUsers)
		api.GET("/books", controller.GetBooks)
	}

	return r
}
