package main

import (
	"fmt"

	"net/http"
	// datagen "github.com/TudorHulban/DataGenerator"
)

func main() {
	http.HandleFunc("/raw", chartRaw)
	http.HandleFunc("/simple", chartSimple)
	http.HandleFunc("/", indexHandler)

	http.ListenAndServe((fmt.Sprintf(":%v", port)), nil)
}

func NewChart(labels, data []float64) *ChartConfig {
	// https://tobiasahlin.com/blog/chartjs-charts-to-get-you-started/

	return &ChartConfig{
		Type: "line",
		Options: ChartOptions{
			Responsive: true,
			Title: ChartTitle{
				Display: true,
				Text:    "Chart 1",
			},
		},
		Data: ChartData{
			Labels: labels,
			Datasets: []ChartDataset{
				ChartDataset{
					Label:       "xxx",
					Fill:        false,
					BorderColor: "#3e95cd",
					Data:        data,
				},
			},
		},
	}
}
