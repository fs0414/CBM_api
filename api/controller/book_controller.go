package controller

import (
	"github.com/soramar/CBM_api/api/repository"
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
)

type BookResponse struct {
	ID    		uint    `json:"id"`
	Title  		string 	`json:"title"`
	ImageUrl 	string 	`json:"image_url"`
	Loanable 	bool 		`json:"loanable"`
	User    struct {
		ID   uint    `json:"id"`
		Name string `json:"name"`
	} `json:"User"`
}

func GetBooks(context *gin.Context) {
	books, err := repository.GetAllBooks()
	fmt.Println("books", err)
	var responseBooks []BookResponse
	for _, book := range books {
		responseUser := BookResponse{
				ID:       book.ID,
				Title:    book.Title,
				ImageUrl: book.ImageUrl,
				Loanable: book.Loanable,
				User: struct {
					ID   uint   `json:"id"`
					Name string `json:"name"`
				}{
					ID:   book.User.ID,
					Name: book.User.Name,
				},
		}
		responseBooks = append(responseBooks, responseUser)
	}
	context.JSON(http.StatusOK, gin.H{"books": responseBooks})
}
