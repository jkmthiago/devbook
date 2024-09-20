package models

import (
	"errors"
	"strings"
	"time"
)

type Post struct {
	Id         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	Autor_id   uint64    `json:"autor_id,omitempty"`
	Autor_nick string    `json:"autor_nick,omitempty"`
	Likes      uint64    `json:"likes"`
	Created_in time.Time `json:"created_in,omitempty"`
}

// Ensures the user is being registered correctly
func (post *Post) Prepare() error {
	if err := post.validate(); err != nil {
		return err
	}

	if err := post.format(); err != nil {
		return err
	}

	return nil
}

// Validates if some of the values is missing
func (post *Post) validate() error {
	if post.Title == "" {
		return errors.New("the title is required and cannot be empty")
	} else if post.Content == "" {
		return errors.New("the content is required and cannot be empty")
	}

	return nil
}

// Formats any empty spaces in the begin ant in the end of the value but not in the middle
func (post *Post) format() error{
	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)

	return nil
}