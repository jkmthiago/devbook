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
	"strconv"
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

	user_id := strconv.FormatUint(userSavedInTheDB.Id, 10)

	answers.JSON(w, http.StatusOK, models.AuthenticationData{Id: user_id, Token: token})
}
