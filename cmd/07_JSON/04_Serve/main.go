package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
)

func DoPOST(toURL string, cfg *GenConfig) error {
	structURL, _ := url.ParseRequestURI(toURL)
	apiURLFormatted := structURL.String()

	var buf bytes.Buffer

	json.NewEncoder(&buf).Encode(cfg)
	req, _ := http.NewRequest("POST", apiURLFormatted, &buf)

	client := &http.Client{}

	res, errDo := client.Do(req)
	if errDo != nil {
		return errDo
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	//log.Println("body:", string(body))

	var data GenData

	errUnmarshal := json.Unmarshal(body, &data)
	if errUnmarshal != nil {
		return errUnmarshal
	}

	log.Println(data.ColumnNames)
	log.Println(data.Rows)

	return nil
}

func main() {
	log.Println(
		"DoPOST: ",
		DoPOST(
			"http://localhost:3002",
			NewConfig(4),
		),
	)
}
