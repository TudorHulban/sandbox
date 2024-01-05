package main

import (
	"bytes"
	"encoding/json"
	"io"

	"net/http"
)

type FieldConfig struct {
	Name         string `json:"name"`
	Type         int    `json:"ftype"`
	Length       int    `json:"flength"`
	PositiveOnly bool   `json:"fpositive"`
	MinValue     int    `json:"min"`
	MaxValue     int    `json:"max"`
}

type GenConfig struct {
	Configuration []FieldConfig `json:"config"`
	NoRows        int64         `json:"norows"`
}

type GenData struct {
	ColumnNames []string        `json:"columns"`
	Rows        [][]interface{} `json:"rowsdata"`
}

func doPost(url string, config *GenConfig) ([][]float64, [][]float64, error) {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(config)

	req, errNew := http.NewRequest("POST", url, buf)
	if errNew != nil {
		return nil, nil, errNew
	}

	var client http.Client

	res, errDo := client.Do(req)
	if errDo != nil {
		return nil, nil, errDo
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	//log.Println("body:", string(body))

	var bodyData GenData

	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return nil, nil, err
	}
	//log.Println(bodyData.ColumnNames)
	//log.Println(bodyData.Rows)

	bodyFloats, err := SliceInterfToFloat(&bodyData.Rows)
	optimPoints := Simplify(*bodyFloats, 2, true)

	return *bodyFloats, optimPoints, err
}
