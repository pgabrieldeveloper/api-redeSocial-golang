package repositorio

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type usuarioRepositorio struct {
	db *sql.DB
}

// NovoRepositorioDeUsuario cria um repositorio de usuarios
func NovoRepositorioDeUsuario(db *sql.DB) *usuarioRepositorio {
	return &usuarioRepositorio{db}
}

// BuscarUsuarios busca usuarios no banco de dados
func (u usuarioRepositorio) BuscarUsuarios(nomeOuNick string) ([]models.Usuario, error) {
	busca := fmt.Sprintf("%%%s%%", nomeOuNick)

	linhas, err := u.db.Query("select id, name, nick, email, createAt from usuarios where name like ? or nick like ?", busca, busca)
	if err != nil {
		return nil, err
	}
	var usuarios []models.Usuario
	for linhas.Next() {
		var usuario models.Usuario
		if err = linhas.Scan(
			&usuario.ID,
			&usuario.Name,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CreateAt,
		); err != nil {
			return nil, err
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}

// BuscarPorId Busca um usuario pelo ID
func (u usuarioRepositorio) BuscarPorId(ID uint64) (models.Usuario, error) {
	linha, err := u.db.Query("select id, name, nick, email, createAt from usuarios where id = ?", ID)
	if err != nil {
		return models.Usuario{}, err
	}
	defer linha.Close()
	var usuario models.Usuario
	if linha.Next() {
		if err := linha.Scan(&usuario.ID,
			&usuario.Name,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CreateAt); err != nil {
			return models.Usuario{}, err
		}
	}

	return usuario, nil
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

func (u usuarioRepositorio) Atualizar(ID uint64, usuario models.Usuario) error {
	statment, err := u.db.Prepare("update usuarios set name = ?, nick = ?, email = ? where id = ?")
	if err != nil {
		return err
	}
	defer statment.Close()
	_, err = statment.Exec(usuario.Name, usuario.Nick, usuario.Email, ID)
	if err != nil {
		return err
	}

	return nil

}

// Deletar Deleta um usuario do banco de dados
func (u usuarioRepositorio) Deletar(ID uint64) error {
	statment, err := u.db.Prepare("delete from usuarios where id = ?")
	if err != nil {
		return err
	}
	defer statment.Close()
	_, err = statment.Exec(ID)
	if err != nil {
		return err
	}

	return nil

}

// BuscarPorEmail Busca um usuario por email no banco de dados
func (u usuarioRepositorio) BuscarPorEmail(email string) (models.Usuario, error) {
	linha, err := u.db.Query("select password from usuarios where email = ?", email)
	var usuario models.Usuario
	if err != nil {
		return usuario, err
	}

	if linha.Next() {
		if err := linha.Scan(&usuario.Password); err != nil {
			return usuario, err
		}
	}
	return usuario, nil
}
