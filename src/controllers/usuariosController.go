package controllers

import (
	"api/src/Responses"
	"api/src/autentication"
	"api/src/db"
	"api/src/models"
	"api/src/repositorio"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {

	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}
	var usuario models.Usuario
	if err = json.Unmarshal(corpoRequisicao, &usuario); err != nil {
		Responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	if err := usuario.Preparar("cadastrar"); err != nil {
		Responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	db, err := db.Conectar()
	if err != nil {
		Responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	usuarioRepositorio := repositorio.NovoRepositorioDeUsuario(db)
	usuario.ID, err = usuarioRepositorio.Criar(usuario)
	if err != nil {
		Responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	Responses.JSON(w, 201, usuario)
}
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("usuario"))
	db, err := db.Conectar()
	if err != nil {
		Responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	usuarioRepositorio := repositorio.NovoRepositorioDeUsuario(db)
	usuarios, err := usuarioRepositorio.BuscarUsuarios(nameOrNick)
	if err != nil {
		Responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	Responses.JSON(w, http.StatusOK, usuarios)
}
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioId, err := strconv.ParseUint(parametros["id"], 10, 64)
	if err != nil {
		Responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	db, err := db.Conectar()
	if err != nil {
		Responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	usuarioRepositorio := repositorio.NovoRepositorioDeUsuario(db)

	usuario, err := usuarioRepositorio.BuscarPorId(usuarioId)
	if err != nil {
		Responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	Responses.JSON(w, http.StatusOK, usuario)
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioId, err := strconv.ParseUint(parametros["id"], 10, 64)
	if err != nil {
		Responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	usuarioIDJSON, err := autentication.ExtrairUserID(r)
	if err != nil {
		Responses.Erro(w, http.StatusUnauthorized, err)
		return
	}
	if usuarioIDJSON != usuarioId {
		Responses.Erro(w, http.StatusForbidden, errors.New("Usuario Logado difere do usuario que esta sendo atualizado"))
		return
	}
	usuarioJSON, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}
	var usuario models.Usuario
	if err := json.Unmarshal(usuarioJSON, &usuario); err != nil {
		Responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	if err := usuario.Preparar("atualizar"); err != nil {
		Responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	db, err := db.Conectar()
	if err != nil {
		Responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	usuarioRepositorio := repositorio.NovoRepositorioDeUsuario(db)
	usuarioRepositorio.Atualizar(usuarioId, usuario)
	Responses.JSON(w, http.StatusNoContent, nil)
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	usuarioID, err := strconv.ParseUint(parametros["id"], 10, 64)
	if err != nil {
		Responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	db, err := db.Conectar()
	if err != nil {
		Responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	usuarioRepositorio := repositorio.NovoRepositorioDeUsuario(db)

	if err := usuarioRepositorio.Deletar(usuarioID); err != nil {
		Responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	Responses.JSON(w, http.StatusNoContent, nil)

}
