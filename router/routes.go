package router

import (
	"net/http"

	controller "github.com/riazXrazor/github-commit-chart/controllers"
)

// Routers type for router
var Routers = map[string][]Router{
	// group
	"api": {
		{
			path:    "/github/{username}/{repo}",
			method:  http.MethodGet,
			handler: controller.GithubRepoCommitInfo(),
		},
		{
			path:    "/github/check/{username}/{repo}",
			method:  http.MethodGet,
			handler: controller.CheckGithubRepoExists(),
		},
	},
	"chart": {
		{
			path:    "/{username}/{repo}",
			method:  http.MethodGet,
			handler: controller.GenerateCommitChart(),
		},
	},
}
