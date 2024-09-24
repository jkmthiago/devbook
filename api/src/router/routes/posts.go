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
		Uri:                      "/users/{user_id}/posts",
		Method:                   http.MethodGet,
		Funtion:                  controllers.SearchPostsFromUser,
		AuthenticationIsRequired: true,
	},
	{
		Uri:                      "/posts/{post_id}",
		Method:                   http.MethodPut,
		Funtion:                  controllers.UpdatePost,
		AuthenticationIsRequired: true,
	},
	{
		Uri:                      "/posts/{post_id}",
		Method:                   http.MethodDelete,
		Funtion:                  controllers.DeletePost,
		AuthenticationIsRequired: true,
	},
	{
		Uri:                      "/posts/{post_id}/like",
		Method:                   http.MethodPost,
		Funtion:                  controllers.Like,
		AuthenticationIsRequired: true,
	},
	{
		Uri:                      "/posts/{post_id}/unlike",
		Method:                   http.MethodPost,
		Funtion:                  controllers.Unlike,
		AuthenticationIsRequired: true,
	},
}
