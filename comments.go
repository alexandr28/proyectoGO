package main

import (
	"./migration"
	"./routes"
	"flag"
	"fmt"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"./commons"
)

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la migración a la BD")
	flag.IntVar(&commons.Port,"port",8080,"Port for server")
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
		Addr:    fmt.Sprintf(":%d",commons.Port),
		Handler: n,
	}
	log.Printf("start server in http://localhost:%d",commons.Port)
	log.Println(server.ListenAndServe())
	log.Println("Finish Program")
}
