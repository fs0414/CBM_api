package router

import (
	"github.com/gin-gonic/gin"
	"github.com/soramar/CBM_api/api/controller"
	"github.com/soramar/CBM_api/api/middleware"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", controller.SayHello)
	api := r.Group("/api", middleware.JWTAuthMiddleware())
	auth := api.Group("auth")
	{
		auth.POST("/register", controller.Register)
		r.POST("api/login", controller.Login)
		api.POST("/logout", controller.Logout)
		api.GET("/users", controller.GetUsers)
		api.GET("/books", controller.GetBooks)
	}

	return r
}
