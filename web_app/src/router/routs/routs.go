package routs

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI                   string
	Method                string
	Function              func(http.ResponseWriter, *http.Request)
	AuthorizationRequired bool
}

func RouteConfig(router *mux.Router) *mux.Router {
	webAppRoutes := loginRouts
	
	for _, webRoute := range webAppRoutes{
		router.HandleFunc(webRoute.URI, webRoute.Function).Methods(webRoute.Method)
	}

	return router
}