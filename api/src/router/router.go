package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Gera um novo gerenciador de rotas com as rotas configuradas
func GenerateNewRoute() *mux.Router  {
	r := mux.NewRouter()
	return routes.RouteConfig(r)
}