package routs

import (
	"net/http"
	"web_app/src/controllers"
)

var userRoutes = []Route{
	{
		URI:                   "/register",
		Method:                http.MethodGet,
		Function:              controllers.LoadRegisterPage,
		AuthorizationRequired: false,
	},
	{
		URI:                   "/register",
		Method:                http.MethodPost,
		Function:              controllers.RegisterNewUser,
		AuthorizationRequired: false,
	},
}
