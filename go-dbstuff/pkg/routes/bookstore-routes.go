package routes

import (
	"github.com/Wanderer0074348/GoServeIt/go-dbstuff/pkg/controllers"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var RegisterBookStoreRoutes = func(router *mux.Router, db *gorm.DB) {
	bookController := controllers.NewBookController(db)

	router.HandleFunc("/book/", bookController.CreateBook).Methods("POST")
	router.HandleFunc("/book/", bookController.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", bookController.GetBookById).Methods("GET")
	// router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	// router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
