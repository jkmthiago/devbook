package routs

import (
	"net/http"
	"web_app/src/middleware"

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
	webAppRoutes = append(webAppRoutes, userRoutes...)
	webAppRoutes = append(webAppRoutes, homePageRoute)
	webAppRoutes = append(webAppRoutes, postRoutes...)

	for _, webRoute := range webAppRoutes {

		if webRoute.AuthorizationRequired {
			router.HandleFunc(webRoute.URI,
				middleware.Logger(
					middleware.AuthenticateCookiesExistence(webRoute.Function),
				),
			).Methods(webRoute.Method)
		} else {
			router.HandleFunc(webRoute.URI, 
				middleware.Logger(webRoute.Function),
			).Methods(webRoute.Method)
		}
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
