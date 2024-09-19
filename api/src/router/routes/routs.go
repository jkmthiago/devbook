package routes

import (
	"api/src/middlewares"
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
	apiRoutes = append(apiRoutes, loginRoute)
	apiRoutes = append(apiRoutes, postsRouts...)

	for _, route := range apiRoutes {

		if route.AuthenticationIsRequired {
			r.HandleFunc(route.Uri, middlewares.Logger(middlewares.Authenticate(route.Funtion))).Methods(route.Method)
		} else {
			r.HandleFunc(route.Uri, middlewares.Logger(route.Funtion)).Methods(route.Method)
		}

		r.HandleFunc(route.Uri, route.Funtion).Methods(route.Method)
	}

	return r
}
