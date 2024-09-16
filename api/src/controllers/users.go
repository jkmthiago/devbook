package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repositorys"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CRUD

// CREATE - POST
// Cria um novo usuário no banco de dados
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating a new user"))

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Error reading request body"))
		return
	}

	var user models.User

	if err = json.Unmarshal(requestBody, &user); err != nil {
		w.Write([]byte("Error converting request to struct"))
		return
	}

	bd, err := db.LoadDataBase()
	if err != nil {
		w.Write([]byte("Error connecting to the database, please check if it's online"))
		return
	}

	defer bd.Close()

	repository := repositorys.NewUsersRepository(bd)
	userId, err := repository.CreateUser(user)
	if err != nil {
		w.Write([]byte("Error creating the user, please check the parameters"))
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(fmt.Sprintf("Usuário criado com sucesso\nID: %d", userId)))
}

// READ 01 - GET PER ID
// Busca um usuário específico
func ReadUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching the user indicated"))

	parameters := mux.Vars(r)

	id, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		w.Write([]byte("Error trying to colect the id to be searched and convert to integer"))
		return
	}

	bd, err := db.LoadDataBase()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error trying to connect to Database, please check if it's online!"))
		return
	}

	defer bd.Close()

	repository := repositorys.NewUsersRepository(bd)
	searchedUser, err := repository.ReadUser(id)
	if err != nil {
		w.Write([]byte("Error Getting data of this user"))
		return
	}

	w.WriteHeader(http.StatusAccepted)
	if err := json.NewEncoder(w).Encode(searchedUser); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error creating the struct body into JSON"))
		return
	}
}

// READ 02 - GET THEM ALL
// Busca todos os usuários do banco de dados
func ReadUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching for all users in the database"))

	bd, err := db.LoadDataBase()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error trying to connect to Database, please check if it's online!"))
		return
	}

	defer bd.Close()

	lines, err := bd.Query("select * from users")
	if err != nil {
		w.Write([]byte("Error searching for users, please check if the database is online!"))
		return
	}

	defer lines.Close()

	var users []models.User
	for lines.Next() {
		var user models.User

		if err := lines.Scan(&user.Id, &user.Name, &user.Nick, &user.Email, &user.Password, &user.CreatedIn); err != nil {
			w.Write([]byte("Error creating users body"))
			return
		}

		users = append(users, user)
	}

	w.WriteHeader(http.StatusAccepted)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.Write([]byte("Error creating the struct body into JSON"))
		return
	}
}

// UPDATE - PUT INSIDE NEW TOYS AND TAKE OUT THE BROKEN ONES
// Atualiza um usuário específico no banco de dados
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating a especific user"))

	parameters := mux.Vars(r)

	id, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		w.Write([]byte("Error converting the id to integer"))
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Error reading the request body, please review the request body!"))
		return
	}

	var user models.User
	if err := json.Unmarshal(requestBody, &user); err != nil {
		w.Write([]byte("Error converting the request body into struct"))
		return
	}

	bd, err := db.LoadDataBase()
	if err != nil {
		w.Write([]byte("Error trying t connect to Database, please check if is's online"))
		return
	}

	defer bd.Close()

	statement, err := bd.Prepare("update user set name = $1, nick = $2, email = $3, password = $4 where id = $5")
	if err != nil {
		w.Write([]byte("Error creating the statement"))
		return
	}

	defer statement.Close()

	if _, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password, id); err != nil {
		w.Write([]byte("Error updating user"))
		return
	}

	w.Write([]byte("User updated successfully"))
}

// DELETE - DELETE AND KILL THAT MOT********R
// Deleta um usuário específico no banco de dados
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting a especific user"))

	parameters := mux.Vars(r)

	id, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		w.Write([]byte("Error converting id to integer!"))
		return
	}

	bd, err := db.LoadDataBase()
	if err != nil {
		w.Write([]byte("Error trying to connecto to the database, please check if it's online!"))
		return
	}

	defer bd.Close()

	statement, err := bd.Prepare("delete from users where id = $1")
	if err != nil {
		w.Write([]byte("Error creating the statement"))
		return
	}

	defer statement.Close()

	if _, err := statement.Exec(id); err != nil {
		w.Write([]byte("Erro trying to delete the user"))
		return
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("User deleted succesfully"))
}
