package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"web_app/src/answers"
	"web_app/src/config"
	"web_app/src/requests"

	"github.com/gorilla/mux"
)

// Calls Api to post a new Post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	post, err := json.Marshal(map[string]string{
		"title":   r.FormValue("title"),
		"content": r.FormValue("content"),
	})

	if err != nil {
		answers.JSON(w, http.StatusBadRequest, answers.Error{Erro: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts", config.ApiURL)

	response, err := requests.RequestAuthenticated(r, http.MethodPost, url, bytes.NewBuffer(post))
	if err != nil {
		answers.JSON(w, http.StatusInternalServerError, answers.Error{Erro: err.Error()})
		return
	}

	if response.StatusCode >= 400 {
		answers.Erro(w, *response)
		return
	}

	answers.JSON(w, response.StatusCode, nil)
}

func Like(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	post_id, err := strconv.ParseUint(parameters["post_id"], 10, 64)
	if err != nil {
		answers.JSON(w, http.StatusBadRequest, answers.Error{Erro: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts/%d/like", config.ApiURL, post_id)
	response, err := requests.RequestAuthenticated(r, http.MethodPost, url, nil)
	if err != nil {
		answers.JSON(w, http.StatusUnprocessableEntity, answers.Error{Erro: err.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		answers.Erro(w, *response)
		return
	}

	answers.JSON(w, response.StatusCode, nil)
}

func Unlike(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	post_id, err := strconv.ParseUint(parameters["post_id"], 10, 64)
	if err != nil {
		answers.JSON(w, http.StatusBadRequest, answers.Error{Erro: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts/%d/unlike", config.ApiURL, post_id)
	response, err := requests.RequestAuthenticated(r, http.MethodPost, url, nil)
	if err != nil {
		answers.JSON(w, http.StatusUnprocessableEntity, answers.Error{Erro: err.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		answers.Erro(w, *response)
		return
	}

	answers.JSON(w, response.StatusCode, nil)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	post_id, err := strconv.ParseUint(parameters["post_id"], 10, 64)
	if err != nil {
		answers.JSON(w, http.StatusBadRequest, answers.Error{Erro: err.Error()})
		return
	}

	r.ParseForm()
	post, err := json.Marshal(map[string]string{
		"title":   r.FormValue("title"),
		"content": r.FormValue("content"),
	})

	if err != nil {
		answers.JSON(w, http.StatusBadRequest, answers.Error{Erro: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts/%d", config.ApiURL, post_id)

	response, err := requests.RequestAuthenticated(r, http.MethodPut, url, bytes.NewBuffer(post))
	if err != nil {
		answers.JSON(w, http.StatusUnprocessableEntity, answers.Error{Erro: err.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		answers.Erro(w, *response)
		return
	}
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	post_id, err := strconv.ParseUint(parameters["post_id"], 10, 64)
	if err != nil {
		answers.JSON(w, http.StatusBadRequest, answers.Error{Erro: err.Error()})
		fmt.Println(err)
		return
	}

	fmt.Println(post_id)

	url := fmt.Sprintf("%s/posts/%d", config.ApiURL, post_id)

	response, err := requests.RequestAuthenticated(r, http.MethodDelete, url, nil)
	if err != nil {
		answers.JSON(w, http.StatusUnprocessableEntity, answers.Error{Erro: err.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		answers.Erro(w, *response)
		return
	}
}