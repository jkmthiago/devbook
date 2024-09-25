package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"web_app/src/answers"
	"web_app/src/config"
	"web_app/src/cookies"
	"web_app/src/requests"

	"github.com/gorilla/mux"
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
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		answers.Erro(w, *response)
		return
	}

	fmt.Println(response.StatusCode)
	answers.JSON(w, response.StatusCode, "{}")
}

func Unfollow(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	user_id, err := strconv.ParseUint(parameters["user_id"], 10, 64)
	if err != nil {
		answers.JSON(w, http.StatusBadRequest, answers.Error{Erro: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/users/%d/unfollow", config.ApiURL, user_id)

	response, err := requests.RequestAuthenticated(r, http.MethodPost, url, nil)
	if err != nil {
		answers.JSON(w, http.StatusInternalServerError, answers.Error{Erro: err.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		answers.Erro(w, *response)
		return
	}

	answers.JSON(w, response.StatusCode, nil)
}

func Follow(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	user_id, err := strconv.ParseUint(parameters["user_id"], 10, 64)
	if err != nil {
		answers.JSON(w, http.StatusBadRequest, answers.Error{Erro: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/users/%d/follow", config.ApiURL, user_id)

	response, err := requests.RequestAuthenticated(r, http.MethodPost, url, nil)
	if err != nil {
		answers.JSON(w, http.StatusInternalServerError, answers.Error{Erro: err.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		answers.Erro(w, *response)
		return
	}

	answers.JSON(w, response.StatusCode, nil)
}

func EditUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user, err := json.Marshal(map[string]string{
		"name":  r.FormValue("name"),
		"nick":  r.FormValue("nick"),
		"email": r.FormValue("email"),
	})
	if err != nil {
		answers.JSON(w, http.StatusBadRequest, answers.Error{Erro: err.Error()})
		return
	}

	cookie, _ := cookies.ReadCookies(r)
	user_id, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/users/%d", config.ApiURL, user_id)

	response, err := requests.RequestAuthenticated(r, http.MethodPut, url, bytes.NewBuffer(user))
	if err != nil {
		answers.JSON(w, http.StatusInternalServerError, answers.Error{Erro: err.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		answers.Erro(w, *response)
		return
	}

	answers.JSON(w, response.StatusCode, nil)
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	passwords, err := json.Marshal(map[string]string{
		"new_Password": r.FormValue("new_Password"),
		"old_Password": r.FormValue("old_Password"),
	})
	if err != nil {
		answers.JSON(w, http.StatusBadRequest, answers.Error{Erro: err.Error()})
		return
	}

	cookie, _ := cookies.ReadCookies(r)
	user_id, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/users/%d/updatePassword", config.ApiURL, user_id)

	response, err := requests.RequestAuthenticated(r, http.MethodPost, url, bytes.NewBuffer(passwords))
	if err != nil {
		answers.JSON(w, http.StatusInternalServerError, answers.Error{Erro: err.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		answers.Erro(w, *response)
		return
	}

	answers.JSON(w, response.StatusCode, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.ReadCookies(r)
	user_id, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/users/%d", config.ApiURL, user_id)
	fmt.Println(url)
	response, err := requests.RequestAuthenticated(r, http.MethodDelete, url, nil)
	if err != nil {
		answers.JSON(w, http.StatusInternalServerError, answers.Error{Erro: err.Error()})
		return
	}

	if response.StatusCode >= 400 {
		answers.Erro(w, *response)
		return		
	}

	answers.JSON(w, http.StatusOK, nil)
}
