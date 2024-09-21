package middleware

import (
	"log"
	"net/http"
	"web_app/src/cookies"
)

func Logger(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunction(w, r)
	}
}

func AuthenticateCookiesExistence(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := cookies.ReadCookies(r); err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		nextFunction(w, r)
	}
}
