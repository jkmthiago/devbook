package repositorys

import (
	"api/src/models"
	"database/sql"
)

type Users struct {
	database *sql.DB
}

func NewUsersRepository(database *sql.DB) *Users {
	return &Users{database}
}

// CREATE - POST
// Cria um novo usuário no banco de dados
func (repository Users) CreateUser(user models.User) (uint64, error) {
	statement, err := repository.database.Prepare("insert into users (name, nick, email, password) values ($1, $2, $3, $4) returning id")
	if err != nil {
		return 0, err
	}

	defer statement.Close()

	var id int64
	err = statement.QueryRow(user.Name, user.Nick, user.Email, user.Password).Scan(&id)
	if err != nil {
		return 0, err
	}

	return uint64(id), nil
}

func (repository Users) ReadUser(id uint64) (models.User, error) {
	var searchedUser models.User

	line, err := repository.database.Query("select * from users where id = $1", id)
	if err != nil {
		return searchedUser, err
	}

	defer line.Close()

	if line.Next() {
		if err := line.Scan(&searchedUser.Id, &searchedUser.Name, &searchedUser.Nick, &searchedUser.Email, &searchedUser.Password, &searchedUser.CreatedIn); err != nil {
			return searchedUser, err
		}
	}

	return searchedUser, nil
}
