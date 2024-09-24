package models

import "time"
 
type Post struct {
	Id         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	Autor_id   uint64    `json:"autor_id,omitempty"`
	Autor_nick string    `json:"autor_nick,omitempty"`
	Likes      uint64    `json:"likes"`
	Created_in time.Time `json:"created_in,omitempty"`
}