package models

import (
	"errors"
	"strings"
	"time"
)

//Usuario modele de um usuario na api
type Usuario struct {
	ID       uint64    `json:id,omitempty`
	Name     string    `json:name,omitempty`
	Nick     string    `json:nick,omitempty`
	Email    string    `json:email,omitempty`
	Password string    `json:password,omitempty`
	CreateAt time.Time `json:creatAt,omitempty`
}

//Preparar valida e prepara campos para ser inseridos no banco de dados
func (u *Usuario) Preparar() error {

	if err := u.validar(); err != nil {
		return err
	}
	u.formatar()
	return nil
}

func (u *Usuario) validar() error {
	if u.Name == "" {
		return errors.New("O campo name é um campo obrigatorio e nao pode estar em branco")
	}
	if u.Nick == "" {
		return errors.New("O campo nick é um campo obrigatorio e nao pode estar em branco")
	}
	if u.Email == "" {
		return errors.New("O campo email é um campo obrigatorio e nao pode estar em branco")
	}
	if u.Password == "" {
		return errors.New("O campo password é um campo obrigatorio e nao pode estar em branco")
	}
	return nil
}

func (u *Usuario) formatar() {
	u.Name = strings.TrimSpace(u.Name)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Nick = strings.TrimSpace(u.Nick)
}
