package controllers

import (
	"api/src/Responses"
	"api/src/db"
	"api/src/models"
	"api/src/repositorio"
	"api/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//Login é responsavel por authenticar um usuario na API
func Login(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}
	var usuario models.Usuario
	if err := json.Unmarshal(corpoRequisicao, &usuario); err != nil {
		Responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	db, err := db.Conectar()
	if err := json.Unmarshal(corpoRequisicao, &usuario); err != nil {
		Responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	usuarioRepositorio := repositorio.NovoRepositorioDeUsuario(db)
	usuarioBanco, err := usuarioRepositorio.BuscarPorEmail(usuario.Email)
	if err := json.Unmarshal(corpoRequisicao, &usuario); err != nil {
		Responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	if err := security.CompareHashPassword(usuarioBanco.Password, usuario.Password); err != nil {
		Responses.Erro(w, http.StatusUnauthorized, err)
		return
	}
	w.Write([]byte("Logado com sucesso"))
}
