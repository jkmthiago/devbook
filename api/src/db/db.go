package db

import (
	"api/src/config"
	"database/sql"
	"fmt"
)

func LoadDataBase() (*sql.DB, error){
	
	var dataSourceName = fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable"+
		config.Db_host, config.Db_port, config.Db_user, config.Db_password, config.Db_name)

	fmt.Println("Iniciando comunicação com o banco de dados")

	db, err := sql.Open("postgres", dataSourceName)

	if err != nil {
		fmt.Println("Conexão Falha com o Banco: Erro de Inicio de Conexão")
		return nil, err
	}

	if err := db.Ping(); err != nil {
		fmt.Println("Conexão Falha com o Banco: Erro de Ping")
		db.Close()
		return nil, err
	}

	fmt.Println("Conexão Bem Sucedida com o Banco:")
	return db, nil
}