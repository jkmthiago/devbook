package routs

import (
	"net/http"
	"web_app/src/controllers"
)

var postRoutes = []Route{
	{
		URI: "/posts",
		Method: "POST",
		Function: controllers.CreatePost,
		AuthorizationRequired: true,
	},
	{
		URI: "/posts/{post_id}/like",
		Method: "POST",
		Function: controllers.Like,
		AuthorizationRequired: true,
	},
	{
		URI: "/posts/{post_id}/unlike",
		Method: "POST",
		Function: controllers.Unlike,
		AuthorizationRequired: true,
	},
	{
		URI: "/posts/{post_id}/edit",
		Method: "GET",
		Function: controllers.LoadEditPage,
		AuthorizationRequired: true,
	},
	{
		URI: "/posts/{post_id}",
		Method: "PUT",
		Function: controllers.UpdatePost,
		AuthorizationRequired: true,
	},
	{
		URI: "/posts/{post_id}",
		Method: http.MethodDelete,
		Function: controllers.DeletePost,
		AuthorizationRequired: true,
	},
}
