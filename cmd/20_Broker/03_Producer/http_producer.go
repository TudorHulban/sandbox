package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type jsonAnswer struct {
	Response string `json:"response"`
}

type request struct {
	codeProducer string
	id           int64
	ttl          int64
	payload      []string
}

func registerWithBroker() (string, error) {
	json := `{"code": "` + producerCode + `", "ip":"` + ip + `"}`

	return jsonRequest(brokerAPIurl, routeRegister, json, "POST")
}

// TODO: refactor
func postRequest(req request) (string, error) {
	json := `{"code": "` + req.codeProducer + `", "id":` + strconv.FormatInt(req.id, 10) + `,"ttl":` + strconv.FormatInt(req.ttl, 10) + `,"payload":["` + strings.Join(req.payload[:], "\", \"") + `"]}`

	fmt.Println("JSON:", json)

	return jsonRequest(brokerAPIurl, routeEvents, json, "POST")
}

func readyRequest(req request) bool {
	json := `{"code": "` + req.codeProducer + `", "id":` + strconv.FormatInt(req.id, 10) + `}`
	log.Println("json:", json)

	ready, errStatus := jsonRequest(brokerAPIurl, routeStatus, json, "POST")
	log.Println(ready, errStatus)

	return false
}

func jsonRequest(urlAPI, route, json, methodHTTP string) (string, error) {
	var client http.Client

	jsonString := []byte(json)

	u, _ := url.ParseRequestURI(urlAPI)
	u.Path = route
	apiURLFormatted := u.String()

	request, errNewRequest := http.NewRequest(methodHTTP, apiURLFormatted, bytes.NewBuffer(jsonString))
	if errNewRequest != nil {
		return "",
			errNewRequest
	}

	request.Header.Set("X-Custom-Header", "myvalue")
	request.Header.Set("Content-Type", "application/json")

	response, errNewRequest := client.Do(request)
	if errNewRequest != nil {
		return "",
			errNewRequest
	}

	return readResponse(response)
}

func readResponse(resp *http.Response) (string, error) {
	defer resp.Body.Close()

	body, errRead := io.ReadAll(resp.Body)
	if errRead != nil {
		return "",
			errRead
	}

	var response jsonAnswer

	if errConvert := json.Unmarshal(body, &response); errConvert != nil {
		return "",
			errConvert
	}

	return response.Response,
		nil
}
