package config

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var (
	Db_user     = ""
	Db_host     = ""
	Db_port     = ""
	Db_password = ""
	Db_name     = ""
	Api_port    = ""
)

// Carrega as variáveis de ambiente
func LoadEnvVar() {
	Api_port = os.Getenv("API_PORT")
	fmt.Println(Api_port, "oh yeah")
	Db_port = os.Getenv("DB-PORT")
	fmt.Println(Db_port, "oh yeah")
	Db_user = os.Getenv("DB_USER")
	Db_host = os.Getenv("DB_HOST")
	Db_password = os.Getenv("DB_USER")
	Db_name = os.Getenv("DB_PASS")
}
