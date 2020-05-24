package libs

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// GetRepoCommitData get commit data from username and repo supplied
func GetRepoCommitData(username, repo string) map[string]int {
	TOKEN := os.Getenv("PERSONAL_ACCESS_TOKEN")
	data := make(map[string]int)
	context := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: TOKEN},
	)
	tc := oauth2.NewClient(context, ts)

	client := github.NewClient(tc)
	// options := &github.CommitsListOptions{
	// 	ListOptions: github.ListOptions{PerPage: 50},
	// }

	commitInfo, _, err := client.Repositories.ListCommits(context, username, repo, nil)
	if err != nil {
		fmt.Printf("Problem in getting repository information %v\n", err)
		os.Exit(1)
	}

	for _, commit := range commitInfo {

		day := commit.Commit.Author.Date.Format("02-Jan-2006")

		if _, ok := data[day]; !ok {
			data[day] = 0
		}
		data[day] = data[day] + 1
	}

	return data
}

// CheckGithubRepo check github repo
func CheckGithubRepo(username, repo string) map[string]int {
	TOKEN := os.Getenv("PERSONAL_ACCESS_TOKEN")
	data := make(map[string]int)
	context := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: TOKEN},
	)
	tc := oauth2.NewClient(context, ts)

	client := github.NewClient(tc)

	_, _, err := client.Repositories.ListCommits(context, username, repo, nil)
	if err != nil {
		data["status"] = 404
	}

	data["status"] = 200

	return data
}
