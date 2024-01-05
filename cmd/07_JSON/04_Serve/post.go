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
	structuredURL, _ := url.ParseRequestURI(toURL)

	var buf bytes.Buffer

	if errEncode := json.NewEncoder(&buf).Encode(cfg); errEncode != nil {
		return errEncode
	}

	req, errRequest := http.NewRequest(
		"POST",
		structuredURL.String(),
		&buf,
	)
	if errRequest != nil {
		return errRequest
	}

	var client http.Client

	res, errDo := client.Do(req)
	if errDo != nil {
		return errDo
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var data GenData

	if errUnmarshal := json.Unmarshal(body, &data); errUnmarshal != nil {
		return errUnmarshal
	}

	log.Println(data.ColumnNames)
	log.Println(data.Rows)

	return nil
}
