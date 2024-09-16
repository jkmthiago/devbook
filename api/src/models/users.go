package models

import "time"

// Representa o tipo de usuario que estará utilizando o  
type User struct {
	Id        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedIn time.Time `json:"createdin,omitempty"`
}
