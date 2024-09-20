package controllers

import (
	"api/src/answers"
	"api/src/authentication"
	"api/src/db"
	"api/src/models"
	"api/src/repositorys"
	"api/src/security"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CRUD

// CREATE - POST
// Cria um novo usuário no banco de dados
func CreateUser(w http.ResponseWriter, r *http.Request) {

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		answers.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User

	if err = json.Unmarshal(requestBody, &user); err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare("register"); err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	bd, err := db.LoadDataBase()
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer bd.Close()

	repository := repositorys.NewUsersRepository(bd)
	user.Id, err = repository.CreateUser(user)
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusCreated, user)
}

// READ - GET
// Busca usuários com nome ou nick parecido
func ReadUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower((r.URL.Query().Get("user")))

	bd, err := db.LoadDataBase()
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer bd.Close()

	repository := repositorys.NewUsersRepository(bd)
	searchedUsers, err := repository.ReadUsers(nameOrNick)
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusOK, searchedUsers)
}

// READ 02 - GET
// Busca um usuário específico
func ReadUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userId, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	bd, err := db.LoadDataBase()
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer bd.Close()

	repository := repositorys.NewUsersRepository(bd)
	searchedUser, err := repository.ReadUser(userId)
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusOK, searchedUser)
}

// UPDATE - PUT INSIDE NEW TOYS AND TAKE OUT THE BROKEN ONES
// Atualiza um usuário específico no banco de dados
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	id, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	userIdOnToken, err := authentication.ExtractUserId(r)
	if err != nil {
		answers.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if id != userIdOnToken {
		answers.Erro(w, http.StatusForbidden, errors.New("it is not possible to change anothe user that ir is not yourself"))
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		answers.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err := json.Unmarshal(requestBody, &user); err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err := user.Prepare("edit"); err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	bd, err := db.LoadDataBase()
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer bd.Close()

	repository := repositorys.NewUsersRepository(bd)
	if err := repository.UpdateUser(id, user); err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusNoContent, nil)
}

// DELETE - DELETE AND KILL THAT MOT********R
// Deleta um usuário específico no banco de dados
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	id, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	userIdOnToken, err := authentication.ExtractUserId(r)
	if err != nil {
		answers.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if id != userIdOnToken {
		answers.Erro(w, http.StatusForbidden, errors.New("it is not possible to change anothe user that ir is not yourself"))
		return
	}

	bd, err := db.LoadDataBase()
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer bd.Close()

	repository := repositorys.NewUsersRepository(bd)
	if err = repository.DeleteUser(id); err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusOK, nil)
}

// POST - POST THAT I AM WITH U
// Cria o sequimento de outro usuário
func FollowUser(w http.ResponseWriter, r *http.Request) {
	follower_id, err := authentication.ExtractUserId(r)
	if err != nil {
		answers.Erro(w, http.StatusUnauthorized, err)
		return
	}

	parameters := mux.Vars(r)

	user_id, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	if user_id == follower_id {
		answers.Erro(w, http.StatusForbidden, errors.New("impossible to follow youtself"))
		return
	}

	bd, err := db.LoadDataBase()
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer bd.Close()

	repository := repositorys.NewUsersRepository(bd)
	if err = repository.FollowUser(user_id, follower_id); err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusNoContent, nil)
}

// Remove o seguimento de outro usuário
func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	follower_id, err := authentication.ExtractUserId(r)
	if err != nil {
		answers.Erro(w, http.StatusUnauthorized, err)
		return
	}

	parameters := mux.Vars(r)
	user_id, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	if user_id == follower_id {
		answers.Erro(w, http.StatusForbidden, errors.New("impossible to unfollow youtself"))
		return
	}

	bd, err := db.LoadDataBase()
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	repository := repositorys.NewUsersRepository(bd)
	if err = repository.UnfollowUser(user_id, follower_id); err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusNoContent, nil)
}

// Busca e retorna todos os usuários que seguem outro usuário
func UserFollowers(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	user_id, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	bd, err := db.LoadDataBase()
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	repository := repositorys.NewUsersRepository(bd)
	followers, err := repository.UserFollowers(user_id)
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusAccepted, followers)
}

// Retorna os usuários que quem solicitou está seguindo
func Following(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	follower_id, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	bd, err := db.LoadDataBase()
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	repository := repositorys.NewUsersRepository(bd)
	following, err := repository.Following(follower_id)
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusAccepted, following)
}

// Atualiza a senha de acesso do usuário que solicitou.
func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	userIdOnToken, err := authentication.ExtractUserId(r)
	if err != nil {
		answers.Erro(w, http.StatusUnauthorized, err)
		return
	}

	parameters := mux.Vars(r)

	user_id, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	if user_id != userIdOnToken {
		answers.Erro(w, http.StatusForbidden, errors.New("not authorized to change anothers password"))
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		answers.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var password models.Password
	if err = json.Unmarshal(requestBody, &password); err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	} else if password.New_Password == password.Old_Password {
		answers.Erro(w, http.StatusBadRequest, errors.New("the passwords are the same"))
		return
	}

	bd, err := db.LoadDataBase()
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer bd.Close()

	repository := repositorys.NewUsersRepository(bd)
	passwordOnTheBank, err := repository.SearchPassword(user_id)
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if err := security.VerifyPassword(passwordOnTheBank, password.Old_Password); err != nil {
		answers.Erro(w, http.StatusUnauthorized, err)
		return
	}

	newHashPassword, err := security.Hash(password.New_Password)
	if err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = repository.UpdatePassword(user_id, string(newHashPassword)); err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusNoContent, nil)
}
