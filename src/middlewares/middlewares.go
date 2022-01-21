package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

// Autenticar responsavel por autenticar um usuario
func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("ANILISANDO")
		next(w, r)
	}
}

// Logger Responsavel por printar informações da rota
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}
