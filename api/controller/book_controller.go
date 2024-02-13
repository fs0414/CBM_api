package controller

import (
	"github.com/soramar/CBM_api/api/repository"
	"github.com/gin-gonic/gin"
	"fmt"
)

func GetBooks(context *gin.Context) {
	books, err := repository.GetAllBooks()
	fmt.Println("books", err)
	context.JSON(200, books)
}
