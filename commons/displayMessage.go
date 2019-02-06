package commons

import (
	"../models"
	"encoding/json"
	"log"
	"net/http"
)

// displaymessage devuelve un mensaje al cliente-
func DisplayMessage(w http.ResponseWriter, m models.Message) {
	j, err := json.Marshal(m)
	if err != nil {
		log.Fatalf("error al convertir el mensaje %s", err)
	}
	w.WriteHeader(m.CodState)
	w.Write(j)
}
