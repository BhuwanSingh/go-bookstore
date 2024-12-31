package routes

import (
	"github.com/BhuwanSingh/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookstore = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookbyID).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
