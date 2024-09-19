package routes

import (
	"api/src/controllers"
	"net/http"
)

var postsRouts = []Route{
	{
		Uri:                      "/posts",
		Method:                   http.MethodPost,
		Funtion:                  controllers.CreatePost,
		AuthenticationIsRequired: true,
	},
	{
		Uri:                      "/posts",
		Method:                   http.MethodGet,
		Funtion:                  controllers.SearchPosts,
		AuthenticationIsRequired: true,
	},
	{
		Uri:                      "/posts/{post_id}",
		Method:                   http.MethodGet,
		Funtion:                  controllers.SearchPost,
		AuthenticationIsRequired: true,
	},
	{
		Uri:                      "/posts/{post_id}",
		Method:                   http.MethodPut,
		Funtion:                  controllers.UpdatePost,
		AuthenticationIsRequired: true,
	},
	{
		Uri:                      "/users",
		Method:                   http.MethodDelete,
		Funtion:                  controllers.DeletePost,
		AuthenticationIsRequired: true,
	},
}