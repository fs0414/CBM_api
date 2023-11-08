package router

import (
	"github.com/soramar/CBM_api/api/controller"
  // "gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
    r := gin.Default()

    r.GET("/", controller.SayHello)
    api := r.Group("/api")
    auth := api.Group("auth")
    {
      auth.POST("/register", controller.Register)
      api.POST("/login", controller.Login)
      api.POST("/logout", controller.Logout)
      api.GET("/users", controller.GetUsers)
      api.GET("/books", controller.GetBooks)
    }

  return r
}
