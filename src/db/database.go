package db

import (
	"api/src/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //Driver de conxao mySQL
)

//Conectar cria uma Conexao no banco de dados
func Conectar() (db *sql.DB, err error) {

	db, erro := sql.Open("mysql", config.StringConexaoBanco)
	if erro != nil {
		return nil, erro
	}
	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}
	return db, nil
}
