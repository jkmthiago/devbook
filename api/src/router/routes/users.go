package routes

import (
	"api/src/controllers"
	"net/http"
)

var usersRouts = []Route{
	{
		Uri:                      "/users",
		Method:                   http.MethodPost,
		Funtion:                  controllers.CreateUser,
		AuthenticationIsRequired: false,
	},
	{
		Uri:                      "/users",
		Method:                   http.MethodGet,
		Funtion:                  controllers.ReadUsers,
		AuthenticationIsRequired: false,
	},
	{
		Uri:                      "/users/{userId}",
		Method:                   http.MethodGet,
		Funtion:                  controllers.ReadUser,
		AuthenticationIsRequired: false,
	},
	{
		Uri:                      "/users/{userId}",
		Method:                   http.MethodPut,
		Funtion:                  controllers.UpdateUser,
		AuthenticationIsRequired: false,
	},
	{
		Uri:                      "/users/{userId}",
		Method:                   http.MethodDelete,
		Funtion:                  controllers.DeleteUser,
		AuthenticationIsRequired: false,
	},
}
