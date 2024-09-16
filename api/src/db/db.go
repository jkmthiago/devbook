package db

import (
	"api/src/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func LoadDataBase() (*sql.DB, error) {
	dataSourceName := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Db_host, config.Db_port, config.Db_user, config.Db_password, config.Db_name,
	)

	fmt.Println("Iniciando comunicação com o banco de dados")

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		fmt.Printf("Erro ao iniciar conexão com o banco: %v\n", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		fmt.Printf("Erro ao conectar com o banco (Ping): %v\n", err)
		db.Close()
		return nil, err
	}

	fmt.Println("Conexão bem-sucedida com o banco de dados")
	return db, nil
}
