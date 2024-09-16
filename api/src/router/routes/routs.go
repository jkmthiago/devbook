package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Cria a estrutura a ser seguida das rotas de requisição da API
type Route struct {
	Uri                      string
	Method                   string
	Funtion                  func(http.ResponseWriter, *http.Request)
	AuthenticationIsRequired bool
}

// Configura todas as rotas no Roteador
func RouteConfig(r *mux.Router) *mux.Router {
	apiRoutes := usersRouts

	for _, route := range apiRoutes {
		r.HandleFunc(route.Uri, route.Funtion).Methods(route.Method)
	}

	return r
}
