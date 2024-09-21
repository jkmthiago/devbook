package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"web_app/src/answers"
	"web_app/src/config"
	"web_app/src/cookies"
	"web_app/src/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})

	if err != nil {
		answers.JSON(w, http.StatusBadRequest, answers.Error{Erro: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/login", config.ApiURL)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(user))
	if err != nil {
		answers.JSON(w, http.StatusInternalServerError, answers.Error{Erro: err.Error()})
		return
	}

	defer response.Body.Close()

	fmt.Println(response.StatusCode)

	if response.StatusCode >= 400 {
		answers.Erro(w, *response)
		return
	}

	var authenticationData models.AuthenticationData
	if err = json.NewDecoder(response.Body).Decode(&authenticationData); err != nil {
		answers.JSON(w, http.StatusUnprocessableEntity, answers.Error{Erro: err.Error()})
		return
	}

	if err = cookies.SaveAuthData(w, authenticationData.Id, authenticationData.Token); err != nil {
		answers.JSON(w, http.StatusUnprocessableEntity, answers.Error{Erro: err.Error()})
		return
	}

	answers.JSON(w, http.StatusOK, nil)
}
