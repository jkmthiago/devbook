package controllers

import (
	"api/src/answers"
	"api/src/authentication"
	"api/src/db"
	"api/src/models"
	"api/src/repositorys"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Cria uma nova postagem
func CreatePost(w http.ResponseWriter, r *http.Request) {
	user_id, err := authentication.ExtractUserId(r)
	if err != nil {
		answers.Erro(w, http.StatusUnauthorized, err)
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		answers.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post models.Post
	if err = json.Unmarshal(requestBody, &post); err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	post.Autor_id = user_id

	if err = post.Prepare(); err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	bd, err := db.LoadDataBase()
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	repository := repositorys.NewPostsRepository(bd)
	post.Id, err = repository.CreatePost(post)
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusCreated, post)
}

// Pesquisa todas as postagens dos usuários que o solicitante segue
func SearchPosts(w http.ResponseWriter, r *http.Request) {
	user_id, err := authentication.ExtractUserId(r)
	if err != nil {
		answers.Erro(w, http.StatusUnauthorized, err)
		return
	}

	bd, err := db.LoadDataBase()
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer bd.Close()

	repository := repositorys.NewPostsRepository(bd)
	posts, err := repository.SearchPosts(user_id)
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusOK, posts)
}

// Pesquisa uma Postagem específica
func SearchPost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	post_id, err := strconv.ParseUint(parameters["post_id"], 10, 64)
	if err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	bd, err := db.LoadDataBase()
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	repository := repositorys.NewPostsRepository(bd)
	post, err := repository.SearchPost(post_id)
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusOK, post)
}

func SearchPostsFromUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	user_id, err := strconv.ParseUint(parameters["user_id"], 10, 64)
	if err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	bd, err := db.LoadDataBase()
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	repository := repositorys.NewPostsRepository(bd)
	posts, err := repository.SearchPostsFromUser(user_id)
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusOK, posts)
}

// Equita/Atualiza uma postagem
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	userIdOnToken, err := authentication.ExtractUserId(r)
	if err != nil {
		answers.Erro(w, http.StatusUnauthorized, err)
		return
	}
	parameters := mux.Vars(r)

	post_id, err := strconv.ParseUint(parameters["post_id"], 10, 64)
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

	repository := repositorys.NewPostsRepository(bd)
	postSavedInTheDB, err := repository.SearchPost(post_id)
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if postSavedInTheDB.Autor_id != userIdOnToken {
		answers.Erro(w, http.StatusForbidden, errors.New("it is forbidden to change others posts"))
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		answers.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post models.Post
	if err = json.Unmarshal(requestBody, &post); err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = post.Prepare(); err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = repository.UpdatePost(post_id, post); err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusNoContent, nil)
}

// Deleta uma postagem
func DeletePost(w http.ResponseWriter, r *http.Request) {
	userIdOnToken, err := authentication.ExtractUserId(r)
	if err != nil {
		answers.Erro(w, http.StatusUnauthorized, err)
		return
	}

	parameters := mux.Vars(r)
	post_id, err := strconv.ParseUint(parameters["post_id"], 10, 64)
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

	repository := repositorys.NewPostsRepository(bd)
	postOnDb, err := repository.SearchPost(post_id)
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if userIdOnToken != postOnDb.Autor_id {
		answers.Erro(w, http.StatusForbidden, errors.New("it is forbidden to delete others posts"))
		return
	}

	if err = repository.DeletePost(post_id); err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusNoContent, nil)
}

func Like(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	post_id, err := strconv.ParseUint(parameters["post_id"], 10, 64)
	if err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	bd, err := db.LoadDataBase()
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	repository := repositorys.NewPostsRepository(bd)
	if err := repository.Like(post_id); err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusNoContent, nil)
}

func Unlike(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	post_id, err := strconv.ParseUint(parameters["post_id"], 10, 64)
	if err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	bd, err := db.LoadDataBase()
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	repository := repositorys.NewPostsRepository(bd)
	if err := repository.Unlike(post_id); err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusNoContent, nil)
}
