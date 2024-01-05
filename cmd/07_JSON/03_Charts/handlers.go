package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func chartRaw(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	dataConfig, _ := datagen.GetConfig(jsonConfig)
	dataChart, _ := datagen.GetData(dataConfig)
	dataFloat, _ := SliceInterfToFloat(&dataChart.Rows)

	log.Println("chartRaw | data:", dataFloat)

	simpleData := Simplify(*dataFloat, 3, false)
	_, simplexData, _ = SplitPointsSlice(simpleData)

	log.Println("raw:", dataFloat)

	var data []float64

	chartLabels, data, _ = SplitPointsSlice(*dataFloat)
	response := NewChart(chartLabels, data)

	json2stream := json.NewEncoder(w)
	json2stream.Encode(&response)

	log.Println("y:", time.Now().Sub(startTime))
}

func chartSimple(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	response := NewChart(chartLabels, simplexData)
	json2stream := json.NewEncoder(w)
	json2stream.Encode(&response)

	log.Println("y:", time.Now().Sub(startTime))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}
