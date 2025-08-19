package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Wanderer0074348/GoServeIt/go-dbstuff/pkg/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type BookController struct {
	DB *gorm.DB
}

func NewBookController(db *gorm.DB) *BookController {
	return &BookController{db}
}

func (c *BookController) CreateBook(w http.ResponseWriter, r *http.Request) {
	book := models.Book{}
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		panic(err)
	}
	if err := c.DB.Create(&book).Error; err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)

}

func (c *BookController) GetBook(w http.ResponseWriter, r *http.Request) {
	book := []models.Book{}
	result := c.DB.Find(&book)
	if result.Error != nil {
		panic(result.Error)
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(book); err != nil {
		panic(err)
	}
}

func (c *BookController) GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	book := models.Book{}
	err := c.DB.First(&book, params["bookId"]).Error

	if err != nil {
		http.Error(w, "Requested book not found", http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(book); err != nil {
		panic(err)
	}
}

// func UpdateBook(w http.ResponseWriter, r *http.Request) {
// 	return
// }

// func DeleteBook(w http.ResponseWriter, r *http.Request) {
// 	return
// }
