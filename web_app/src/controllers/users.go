package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"web_app/src/answers"
	"web_app/src/config"
)

func RegisterNewUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"name":     r.FormValue("name"),
		"nick":     r.FormValue("nick"),
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})

	if err != nil {
		answers.JSON(w, http.StatusBadRequest, answers.Error{Erro: err.Error()})
	}

	url := fmt.Sprintf("%s/users", config.ApiURL)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(user))
	if err != nil {
		answers.JSON(w, http.StatusInternalServerError, answers.Error{Erro: err.Error()})
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		answers.Erro(w, *response)
		return
	}

	fmt.Println(response.StatusCode)
	answers.JSON(w, response.StatusCode, "{}")
}
