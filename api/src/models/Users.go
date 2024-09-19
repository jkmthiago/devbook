package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// Representa o tipo de usuario que estará utilizando o
type User struct {
	Id        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedIn time.Time `json:"createdin,omitempty"`
}

// Ensures the user is being registered correctly
func (user *User) Prepare(stage string) error {
	if err := user.validate(stage); err != nil {
		return err
	}

	if err := user.format(stage); err != nil {
		return err
	}

	return nil
}

// Validates if some of the values is missing
func (user *User) validate(stage string) error {
	if user.Name == "" {
		return errors.New("the name is required and cannot be empty")
	} else if user.Nick == "" {
		return errors.New("the nick is required and cannot be empty")
	} else if user.Email == "" {
		return errors.New("the email is required and cannot be empty")
	} else if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("the email is not valid")
	} else if stage == "register" && user.Password == "" {
		return errors.New("the password is required and cannot be empty")
	}
	return nil
}

// Formats any empty spaces in the begin ant in the end of the value but not in the middle
func (user *User) format(stage string) error{
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if stage == "register"{
		hashPassword, err := security.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(hashPassword)
	}

	return nil
}