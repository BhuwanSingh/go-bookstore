package routes

import (
	"github/gorilla/mux"
)

var RegisterBookstore = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.createBook).Methods
}
