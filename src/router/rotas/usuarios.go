package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI:               "/usuarios",
		Metodo:            http.MethodPost,
		Funcao:            controllers.CriarUsuario,
		ExigeAutenticacao: false,
	},
	{
		URI:               "/usuarios",
		Metodo:            http.MethodGet,
		Funcao:            controllers.BuscarUsuarios,
		ExigeAutenticacao: true,
	},
	{
		URI:               "/usuarios/{id}",
		Metodo:            http.MethodGet,
		Funcao:            controllers.BuscarUsuario,
		ExigeAutenticacao: true,
	},
	{
		URI:               "/usuarios/{id}",
		Metodo:            http.MethodPut,
		Funcao:            controllers.AtualizarUsuario,
		ExigeAutenticacao: false,
	},
	{
		URI:               "/usuarios/{id}",
		Metodo:            http.MethodDelete,
		Funcao:            controllers.DeletarUsuario,
		ExigeAutenticacao: false,
	},
}
