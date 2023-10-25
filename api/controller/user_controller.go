package controller

import (
	"github.com/soramar/CBM_api/api/repository"
	"github.com/gin-gonic/gin"
)

func GetUsers(context *gin.Context) {
	users := repository.GetAll() 
	context.JSON(200, users)
}
