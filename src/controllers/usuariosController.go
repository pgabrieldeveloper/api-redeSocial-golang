package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repositorio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {

	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var usuario models.Usuario
	if err = json.Unmarshal(corpoRequisicao, &usuario); err != nil {
		log.Fatal(err)
	}

	db, err := db.Conectar()
	if err != nil {
		log.Fatal(err)
	}
	usuarioRepositorio := repositorio.NovoRepositorioDeUsuario(db)
	usuarioID, err := usuarioRepositorio.Criar(usuario)
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(fmt.Sprintf("ultimo id inserido: %d", usuarioID)))
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
