package main

import (
	"./migration"
	"./routes"
	"flag"
	"github.com/urfave/negroni"
	"log"
	"net/http"
)

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la migración a la BD")
	flag.Parse()
	if migrate == "yes" {
		log.Println("Comenzo la Migración....")
		migration.Migrate()
		log.Println("Finalizo la Migración")
	}
	// inicia las rutas
	router := routes.InitRoutes()
	// inicia los middlewares
	n := negroni.Classic()
	n.UseHandler(router)
	// iniciamos el servidor
	server := &http.Server{
		Addr:    ":8080",
		Handler: n,
	}
	log.Println("start server in http://localhost:8080")
	log.Println(server.ListenAndServe())
	log.Println("Finish Program")
}
