package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type jsonAnswer struct {
	Response string `json:"response"`
}

func registerWithBroker() (string, error) {
	json := `{"code": "` + consumerCode + `", "socket":"` + socket + `"}` // to change.
	//log.Println("json:", json)

	return jsonRequest(brokerAPIurl, routeRegister, json, "POST")
}

func consumeEvent() (string, error) {
	json := `{"code": "` + consumerCode + `"}` // to change.
	//log.Println("json:", json)
	return jsonRequest(brokerAPIurl, routeConsume, json, "POST")
}

func jsonRequest(urlAPI, route, rawJSON, methodHTTP string) (string, error) {
	var client http.Client
	jsonString := []byte(rawJSON)

	u, _ := url.ParseRequestURI(urlAPI)
	u.Path = route
	apiURLFormatted := u.String()

	request, err := http.NewRequest(methodHTTP, apiURLFormatted, bytes.NewBuffer(jsonString))
	request.Header.Set("X-Custom-Header", "myvalue")
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Println("1")

		return "", err
	}

	response, err := client.Do(request)
	if err != nil {
		log.Println("2")

		return "", err
	}

	return readResponse(response)
}

func readResponse(response *http.Response) (string, error) {
	body, errRead := ioutil.ReadAll(response.Body)
	if errRead != nil {
		return "", errRead
	}
	defer response.Body.Close()

	var result jsonAnswer
	errConvert := json.Unmarshal(body, &result)

	return result.Response, errConvert
}
