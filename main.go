package main

//go:generate go run main.go
//148afcc15f7feb292e66bd55a3bd943d11b05d71
import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"github.com/wcharczuk/go-chart"
	"golang.org/x/oauth2"
)

// ListOptions specifies the optional parameters to various List methods that
// support pagination.
type ListOptions struct {
	// For paginated result sets, the number of results to include per page.
	PerPage int `url:"per_page,omitempty"`
}

func main() {
	//ngx-lottie/ngx-lottie
	username := "ngx-lottie"
	repo := "ngx-lottie"
	var bars []chart.Value
	data := getRepoCommitData(username, repo)

	for key, value := range data {

		record := chart.Value{
			Value: float64(value),
			Label: key,
		}
		bars = append(bars, record)
	}

	graph := chart.BarChart{
		Title: "Commit activity for " + username + "/" + repo,
		Background: chart.Style{
			Padding: chart.Box{
				Top:   40,
				Right: 40,
			},
		},
		XAxis: chart.Style{
			StrokeWidth: 1,
			FontSize:    5,
		},
		YAxis: chart.YAxis{
			ValueFormatter: func(v interface{}) string {
				if vf, isFloat := v.(float64); isFloat {
					return fmt.Sprintf("%0.0f", vf)
				}
				return ""
			},
			Style: chart.Style{
				StrokeWidth: 1,
				FontSize:    5,
			},
		},
		Height:   400,
		BarWidth: 50,
		Bars:     bars,
	}

	f, _ := os.Create("output.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}

func getRepoCommitData(username string, repo string) map[string]int {

	data := make(map[string]int)
	context := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "148afcc15f7feb292e66bd55a3bd943d11b05d71"},
	)
	tc := oauth2.NewClient(context, ts)

	client := github.NewClient(tc)
	options := &github.CommitsListOptions{
		ListOptions: github.ListOptions{PerPage: 50},
	}

	commitInfo, _, err := client.Repositories.ListCommits(context, username, repo, options)
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
