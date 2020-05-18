package router

import (
	"net/http"

	controller "github.com/riazXrazor/github-commit-chart/controllers"
)

// Routers type for router
var Routers = []Router{
	{
		path:    "/",
		method:  http.MethodGet,
		handler: controller.Index(),
	},
	{
		path:    "/home",
		method:  http.MethodGet,
		handler: controller.Index(),
	},
}
