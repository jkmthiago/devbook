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
	{
		URI:                   "/search-users",
		Method:                http.MethodGet,
		Function:              controllers.LoadSearchedUsersPage,
		AuthorizationRequired: true,
	},
	{
		URI:                   "/users/{user_id}",
		Method:                http.MethodGet,
		Function:              controllers.LoadUserPage,
		AuthorizationRequired: true,
	},
	{
		URI:                   "/users/{user_id}/unfollow",
		Method:                http.MethodPost,
		Function:              controllers.Unfollow,
		AuthorizationRequired: true,
	},
	{
		URI:                   "/users/{user_id}/follow",
		Method:                http.MethodPost,
		Function:              controllers.Follow,
		AuthorizationRequired: true,
	},
	{
		URI:                   "/edit-user",
		Method:                http.MethodGet,
		Function:              controllers.LoadEditUser,
		AuthorizationRequired: true,
	},
	{
		URI:                   "/edit-user",
		Method:                http.MethodPut,
		Function:              controllers.EditUser,
		AuthorizationRequired: true,
	},
	{
		URI:                   "/updatePassword",
		Method:                http.MethodGet,
		Function:              controllers.LoadUpdatePassword,
		AuthorizationRequired: true,
	},
	{
		URI:                   "/updatePassword",
		Method:                http.MethodPost,
		Function:              controllers.UpdatePassword,
		AuthorizationRequired: true,
	},
	{
		URI:                   "/delete-user",
		Method:                http.MethodDelete,
		Function:              controllers.DeleteUser,
		AuthorizationRequired: true,
	},
}
