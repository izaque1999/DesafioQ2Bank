package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"q2bank/models"
)

func List(w http.ResponseWriter, e *http.Request) {
	usuarios, err := models.GetAll()
	if err != nil {
		log.Printf("Erro ao obter registros: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuarios)
}
