package libs

import (
	"fmt"

	"github.com/wcharczuk/go-chart"
)

// GenerateChart generate chart from repo commits
func GenerateChart(data map[string]int, username, repo string) *chart.BarChart {
	var bars []chart.Value
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
			Show:        true,
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
				Show:        true,
				StrokeWidth: 1,
				FontSize:    7,
			},
		},
		Height:   400,
		BarWidth: 25,
		Bars:     bars,
	}

	return &graph
}
