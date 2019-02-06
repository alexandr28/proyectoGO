package routes

import (
	"../controllers"
	"github.com/gorilla/mux"
)

// SetLogin Router  router para login
func SetLoginRouter(router *mux.Router) {
	router.HandleFunc("/api/login", controllers.Login).Methods("POST")
}
