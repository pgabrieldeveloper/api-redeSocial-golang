package controllers

import (
	"api/Responses"
	"api/src/db"
	"api/src/models"
	"api/src/repositorio"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {

	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Responses.Erro(w, http.StatusUnprocessableEntity, err)
	}
	var usuario models.Usuario
	if err = json.Unmarshal(corpoRequisicao, &usuario); err != nil {
		Responses.Erro(w, http.StatusBadRequest, err)
	}

	db, err := db.Conectar()
	if err != nil {
		Responses.Erro(w, http.StatusInternalServerError, err)
	}
	usuarioRepositorio := repositorio.NovoRepositorioDeUsuario(db)
	usuario.ID, err = usuarioRepositorio.Criar(usuario)
	if err != nil {
		Responses.Erro(w, http.StatusInternalServerError, err)
	}
	Responses.JSON(w, 201, usuario)
}
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuarios"))
}
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuario"))
}
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuario"))
}
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletanado usuario"))
}
