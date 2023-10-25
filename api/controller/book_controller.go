package controller

import (
	"github.com/soramar/CBM_api/api/repository"
	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	books := repository.GetAll()
	c.JSON(200, books)
}
