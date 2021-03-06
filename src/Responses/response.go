package Responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON retor uma reposta no formato de json para a requisição
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)
	if dados != nil {
		if err := json.NewEncoder(w).Encode(dados); err != nil {
			log.Fatal(err)
		}
	}
}

// Erro retorna um erro no formato json
func Erro(w http.ResponseWriter, statusCode int, err error) {
	w.Header().Set("Content-type", "application/json")
	JSON(w, statusCode, struct {
		Err string `json:"erro"`
	}{Err: err.Error()})
}
