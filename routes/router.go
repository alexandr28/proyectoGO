package routes

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	SetLoginRouter(router)
	SetUserRouter(router)
	SetCommetRouter(router)
	return router
}
