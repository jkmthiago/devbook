package repositorys

import (
	"database/sql"
	"net/http"
)

type Posts struct {
	database *sql.DB
}

func NewPostsRepository(database *sql.DB) *Posts {
	return &Posts{database}
}

func (repository Posts) CreatePost() error {
	return nil
}
func (repository Posts) SearchPosts() error {
	return nil
}
func (repository Posts) SearchPost() error {
	return nil
}
func (repository Posts) UpdatePost() error {
	return nil
}
func (repository Posts) DeletePost() error {
	return nil
}

