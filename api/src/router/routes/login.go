package routes

import (
	"api/src/controllers"
	"net/http"
)

var loginRoute = Route{
	Uri:                      "/login",
	Method:                   http.MethodPost,
	Funtion:                  controllers.Login,
	AuthenticationIsRequired: false,
}
