package routs

import (
	"net/http"
	"web_app/src/controllers"
)

var homePageRoute = Route{
	URI:                   "/home",
	Method:                http.MethodGet,
	Function:              controllers.LoadHomePage,
	AuthorizationRequired: true,
}
