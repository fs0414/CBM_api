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
    {
      api.GET("/auth", controller.GetUsers)
      api.GET("/users", controller.GetUsers)
      api.GET("/books", controller.GetBooks)
    }

  return r
}
