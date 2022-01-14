package router

import (
	"api/src/router/rotas"
	"github.com/gorilla/mux"
)

//Gerar retornar um router com as rotas configuradas
func Gerar() *mux.Router {
	return rotas.Configurar(mux.NewRouter())
}
