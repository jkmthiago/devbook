package controllers

import (
	"api/src/answers"
	"api/src/authentication"
	"api/src/db"
	"api/src/models"
	"api/src/repositorys"
	"api/src/security"
	"encoding/json"
	"io"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
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

	bd, err := db.LoadDataBase()
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer bd.Close()

	repository := repositorys.NewUsersRepository(bd)
	userSavedInTheDB, err := repository.SearchEmail(user.Email)
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.VerifyPassword(userSavedInTheDB.Password, user.Password); err != nil {
		answers.Erro(w, http.StatusUnauthorized, err)
		return
	}

	token, err := authentication.CreateToken(userSavedInTheDB.Id)
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	w.Write([]byte(token))
}
