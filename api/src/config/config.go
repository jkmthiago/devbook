package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	Db_user     = ""
	Db_host     = ""
	Db_port     = 0
	Db_password = ""
	Db_name     = ""
	Api_port    = 0
)

// Carrega as variáveis de ambiente
func LoadEnvVar() {
	Api_port= os.Getenv("API_PORT")
	Db_port = os.Getenv("DB-PORT")
	Db_user = os.Getenv("DB_USER")
	Db_host = os.Getenv("DB_HOST")
	Db_password = os.Getenv("DB_USER")
	Db_name = os.Getenv("DB_PASS")
}
