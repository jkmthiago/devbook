package routs

import (
	"net/http"
	"web_app/src/controllers"
)

var loginRouts = []Route{
	{
		URI:    "/",
		Method: http.MethodGet,
		Function: controllers.LoadLoginPage,
		AuthorizationRequired: false,
	},
	{
		URI:    "/login",
		Method: http.MethodGet,
		Function: controllers.LoadLoginPage,
		AuthorizationRequired: false,
	},
}
