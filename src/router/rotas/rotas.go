package rotas

import (
	"api/src/middlewares"
	"github.com/gorilla/mux"
	"net/http"
)

//Rota representa todas as rotas da api
type Rota struct {
	URI               string
	Metodo            string
	Funcao            func(http.ResponseWriter, *http.Request)
	ExigeAutenticacao bool
}

//Configura Rotas do router
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios
	rotas = append(rotas, rotaLogin)
	for _, rota := range rotas {
		if rota.ExigeAutenticacao {
			r.HandleFunc(rota.URI, middlewares.Logger(middlewares.Autenticar(rota.Funcao))).Methods(rota.Metodo)
		}
		r.HandleFunc(rota.URI, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
	}
	return r
}
