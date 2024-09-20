package router

import (
	"web_app/src/router/routs"

	"github.com/gorilla/mux"
)

func Generate() *mux.Router {
	r := mux.NewRouter()
	return routs.RouteConfig(r)
}
