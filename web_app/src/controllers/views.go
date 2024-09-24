package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"web_app/src/answers"
	"web_app/src/config"
	"web_app/src/cookies"
	"web_app/src/models"
	"web_app/src/requests"
	"web_app/src/utils"

	"github.com/gorilla/mux"
)

func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.ReadCookies(r)

	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	}

	utils.ExecTemplates(w, "login.html", nil)
}

func LoadRegisterPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecTemplates(w, "register.html", nil)
}

func LoadHomePage(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/posts", config.ApiURL)
	response, err := requests.RequestAuthenticated(r, http.MethodGet, url, nil)
	fmt.Println(response.StatusCode, err)

	if err != nil {
		answers.JSON(w, http.StatusInternalServerError, answers.Error{Erro: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		answers.Erro(w, *response)
		return
	}

	var posts []models.Post
	if err := json.NewDecoder(response.Body).Decode(&posts); err != nil {
		answers.JSON(w, http.StatusUnprocessableEntity, answers.Error{Erro: err.Error()})
		return
	}

	cookie, _ := cookies.ReadCookies(r)
	userId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.ExecTemplates(w, "home.html", struct {
		Posts  []models.Post
		UserId uint64
	}{
		Posts:  posts,
		UserId: userId,
	})
}

func LoadEditPage(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	post_id, err := strconv.ParseUint(parameters["post_id"], 10, 64)
	if err != nil {
		answers.JSON(w, http.StatusBadRequest, answers.Error{Erro: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts/%d", config.ApiURL, post_id)
	response, err := requests.RequestAuthenticated(r, http.MethodGet, url, nil)
	if err != nil {
		answers.JSON(w, http.StatusInternalServerError, answers.Error{Erro: err.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		answers.Erro(w, *response)
		return
	}

	var post models.Post
	if err = json.NewDecoder(response.Body).Decode(&post); err != nil {
		answers.JSON(w, http.StatusUnprocessableEntity, answers.Error{Erro: err.Error()})
		return
	}

	utils.ExecTemplates(w, "updatePost.html", post)
}

func LoadSearchedUsersPage(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("search_users"))

	if nameOrNick == "" {
		answers.JSON(w, http.StatusBadRequest, answers.Error{Erro: "Parâmetro de pesquisa vazio"})
		return
	}

	url := fmt.Sprintf("%s/users?user=%s", config.ApiURL, nameOrNick)

	response, err := requests.RequestAuthenticated(r, http.MethodGet, url, nil)
	if err != nil {
		answers.JSON(w, http.StatusInternalServerError, answers.Error{Erro: err.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		answers.Erro(w, *response)
		return
	}

	var users []models.User
	if err = json.NewDecoder(response.Body).Decode(&users); err != nil {
		answers.JSON(w, http.StatusUnprocessableEntity, answers.Error{Erro: err.Error()})
		return
	}

	utils.ExecTemplates(w, "users.html", users)
}

func LoadUserPage(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	user_id, err := strconv.ParseUint(parameters["user_id"], 10, 64)
	if err != nil {
		answers.JSON(w, http.StatusBadRequest, answers.Error{Erro: err.Error()})
		return
	}

	user, err := models.SearchUsersCompleteData(user_id, r)
	if err != nil {
		answers.JSON(w, http.StatusInternalServerError, answers.Error{Erro: err.Error()})
		return
	}

	cookie, _ := cookies.ReadCookies(r)
	loggedUserId, _ := strconv.ParseUint(cookie["id"], 10, 64)
	fmt.Println(loggedUserId)

	utils.ExecTemplates(w, "user.html", struct {
		User         models.User
		LoggedUserId uint64
	}{
		User:         user,
		LoggedUserId: loggedUserId,
	})
}
