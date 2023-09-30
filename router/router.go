package router

import (
	"github.com/soramar/CBM_api/controller"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
  r := gin.Default()
  r.GET("/", controller.SayHello)

  return r
}