package utils

import (
	"net/http"
	"text/template"
)

var templates *template.Template

func LoadTemplates()  {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

func ExecTemplates(w http.ResponseWriter, templateName string, data interface{})  {
	templates.ExecuteTemplate(w, templateName, data)
}