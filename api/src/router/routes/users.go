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
		AuthenticationIsRequired: true,
	},
	{
		Uri:                      "/users/{id}",
		Method:                   http.MethodGet,
		Funtion:                  controllers.ReadUser,
		AuthenticationIsRequired: true,
	},
	{
		Uri:                      "/users/{id}",
		Method:                   http.MethodPut,
		Funtion:                  controllers.UpdateUser,
		AuthenticationIsRequired: true,
	},
	{
		Uri:                      "/users/{id}",
		Method:                   http.MethodDelete,
		Funtion:                  controllers.DeleteUser,
		AuthenticationIsRequired: true,
	},
	{
		Uri:                      "/users/{id}/follow",
		Method:                   http.MethodPost,
		Funtion:                  controllers.FollowUser,
		AuthenticationIsRequired: true,
	},
	{
		Uri:                      "/users/{id}/unfollow",
		Method:                   http.MethodPost,
		Funtion:                  controllers.UnfollowUser,
		AuthenticationIsRequired: true,
	},
	{
		Uri:                      "/users/{id}/followers",
		Method:                   http.MethodGet,
		Funtion:                  controllers.UserFollowers,
		AuthenticationIsRequired: true,
	},
	{
		Uri:                      "/users/{id}/following",
		Method:                   http.MethodGet,
		Funtion:                  controllers.Following,
		AuthenticationIsRequired: true,
	},
	{
		Uri:                      "/users/{id}/updatePassword",
		Method:                   http.MethodPost,
		Funtion:                  controllers.UpdatePassword,
		AuthenticationIsRequired: true,
	},
}
