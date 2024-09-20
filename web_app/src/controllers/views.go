package controllers

import (
	"net/http"
	"web_app/src/utils"
)

func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecTemplates(w, "login.html", nil)
}

func LoadRegisterPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecTemplates(w, "register.html", nil)
}
