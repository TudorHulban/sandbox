package main

type ChartTitle struct {
	Display bool   `json:"display"`
	Text    string `json:"text"`
}

type ChartOptions struct {
	Responsive bool       `json:"responsive"`
	Title      ChartTitle `json:title`
}

type ChartDataset struct {
	Data        []float64 `json:"data"`
	Label       string    `json:"label"`
	Fill        bool      `json:"fill"`
	BorderColor string    `json:"borderColor"`
}

type ChartData struct {
	Labels   []float64      `json:"labels"`
	Datasets []ChartDataset `json:"datasets"`
}

type ChartConfig struct {
	Type    string       `json:"type"`
	Data    ChartData    `json:"data"`
	Options ChartOptions `json:"options"`
}

//var chartData ChartResponse

// chartLabels - labels common to raw and optimized points
var chartLabels []float64

// simplexData - data points after simplification
var simplexData []float64
