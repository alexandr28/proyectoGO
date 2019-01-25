package main

import (
	"./migration"
	"flag"
	"log"
)

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "yes", "Genera la migración a la BD")
	flag.Parse()
	if migrate == "yes" {
		log.Println("Comenzo la Migración....")
		migration.Migrate()
		log.Println("Finalizo la Migración")
	}
}
