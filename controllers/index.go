package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/riazXrazor/github-commit-chart/libs"
	"github.com/wcharczuk/go-chart"
)

// Index display landing page
func Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
	}
}

// CommitInfo display landing page
type CommitInfo struct {
	Date        string `json:"date"`
	NoOfCommits string `json:"no_of_commits"`
}

// GithubRepoCommitInfo display landing page
func GithubRepoCommitInfo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := mux.Vars(r)["username"]
		repo := mux.Vars(r)["repo"]

		data := libs.GetRepoCommitData(username, repo)
		json, _ := json.Marshal(data)
		fmt.Fprint(w, string(json))
	}
}

// GenerateCommitChart display landing page
func GenerateCommitChart() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := mux.Vars(r)["username"]
		repo := mux.Vars(r)["repo"]
		data := libs.GetRepoCommitData(username, repo)

		w.Header().Add("content-type", "image/svg+xml;charset=utf-8")
		w.Header().Add("cache-control", "public, max-age=86400")
		w.Header().Add("date", time.Now().Format(time.RFC1123))
		w.Header().Add("expires", time.Now().Format(time.RFC1123))
		g := libs.GenerateChart(data, username, repo)

		g.Render(chart.SVG, w)
	}
}

// CheckGithubRepoExists check repo exists
func CheckGithubRepoExists() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := mux.Vars(r)["username"]
		repo := mux.Vars(r)["repo"]

		data := libs.CheckGithubRepo(username, repo)
		json, _ := json.Marshal(data)
		fmt.Fprint(w, string(json))
	}
}
