package repositorio

import (
	"api/src/models"
	"database/sql"
)

type usuarioRepositorio struct {
	db *sql.DB
}

// NovoRepositorioDeUsuario cria um repositorio de usuarios
func NovoRepositorioDeUsuario(db *sql.DB) *usuarioRepositorio {
	return &usuarioRepositorio{db}
}

func (u usuarioRepositorio) Criar(usuario models.Usuario) (uint64, error) {
	statment, err := u.db.Prepare("insert into usuarios (name, nick, password,email) values (?,?,?,?)")
	if err != nil {
		return 0, err
	}
	defer statment.Close()
	resultado, err := statment.Exec(usuario.Name, usuario.Nick, usuario.Password, usuario.Email)
	if err != nil {
		return 0, err
	}
	ultimoIDInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(ultimoIDInserido), nil

}
