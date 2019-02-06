package routes

import (
	"../controllers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func SetCommetRouter(router *mux.Router) {
	prefix := "/api/comments/"
	subrouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)

	subrouter.HandleFunc("/", controllers.CommentCreate).Methods("POST")
	subrouter.HandleFunc("/", controllers.CommentGetAll).Methods("GET")
	router.PathPrefix(prefix).Handler(negroni.New(negroni.HandlerFunc(controllers.ValidateToken), negroni.Wrap(subrouter)))
}
