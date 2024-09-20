package controllers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"web_app/src/answers"
)

func RegisterNewUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"nome":     r.FormValue("name"),
		"nicl":     r.FormValue("nick"),
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})

	if err != nil {
		answers.JSON(w, http.StatusBadRequest, answers.Error{Erro: err.Error()})
	}

	response, err := http.Post("http://localhost:5000/users", "application/json", bytes.NewBuffer(user))
	if err != nil {
		answers.JSON(w, http.StatusInternalServerError, answers.Error{Erro: err.Error()})
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		answers.Erro(w, *response)
		return
	}

	answers.JSON(w, response.StatusCode, nil)
}
