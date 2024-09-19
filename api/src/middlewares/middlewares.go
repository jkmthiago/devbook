package middlewares

import (
	"api/src/answers"
	"api/src/authentication"
	"fmt"
	"log"
	"net/http"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n%s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Authenticating...")
		if err := authentication.ValidateToken(r); err != nil{
			answers.Erro(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}