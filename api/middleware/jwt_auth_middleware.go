package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/soramar/CBM_api/api/repository"
	"net/http"
	"os"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
			return
		}

		tokenString = tokenString[len("Bearer "):] // Remove "Bearer " prefix

		if isTokenInvalid := repository.IsTokenInvalid(tokenString); isTokenInvalid {
			fmt.Println("repository.IsTokenInvalid after")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "The token is already signed out"})
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
			user_id, ok_user := claims["user_id"].(string)
			
			if !ok && !ok_user {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
				return
			}

			if repository.IsEmailUnregistered(email) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not registered"})
				return
			}

			c.Set("claims", claims)
			c.Set("tokenString", tokenString)
			c.Set("user_id", user_id)
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		}
	}
}
