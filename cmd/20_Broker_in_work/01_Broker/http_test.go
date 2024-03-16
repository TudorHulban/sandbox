package main

import (
	"bytes"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

const apiURL = "http://localhost:" + port

func TestRegisterProducer(t *testing.T) {
	var client http.Client

	jsonString := []byte(`{"code": "P1", "ip": "127.0.0.5"}`)

	//curl -X POST -H 'Content-Type: application/json' -d "{\"code\": \"P77\", \"ip\": \"127.0.0.77\"}" http://127.0.0.1:8080/registerproducer
	//apiURL := "http://requestbin.fullcontact.com"
	//route := "/yx8kxbyx/"

	u, _ := url.ParseRequestURI(apiURL)
	u.Path = urlRegisterProducer

	request, errReq := http.NewRequest("POST", u.String(), bytes.NewBuffer(jsonString))
	if errReq != nil {
		t.Error("NewRequest constructor error: ", errReq)
	}

	request.Header.Set("X-Custom-Header", "myvalue")
	request.Header.Set("Content-Type", "application/json")

	response, errDo := client.Do(request)
	if errDo != nil {
		t.Error("NewRequest  error: ", errDo)
	}

	r, errRead := readResponse(response)
	require.NoError(t, errRead)
	if r != "ok" {
		t.Error("producer not registered correctly", r)
	}
}

func TestPostEvent(t *testing.T) {
	var client http.Client

	jsonString := []byte(`{"code": "P1", "id": 1001, "ttl": 100, "payload": ["url1", "url2"]}`)

	u, _ := url.ParseRequestURI(apiURL)
	u.Path = urlPostEvents

	request, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(jsonString))
	if err != nil {
		t.Error("NewRequest constructor error: ", err)
	}

	request.Header.Set("X-Custom-Header", "myvalue")
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		t.Error("NewRequest  error: ", err)
	}

	r, _ := readResponse(response)
	if r != "ok" {
		t.Error("producer not registered correctly", r)
	}
}

func TestConsumeEvent(t *testing.T) {
	var client http.Client

	jsonString := []byte(`{"code": "C1"}`)

	u, _ := url.ParseRequestURI(apiURL)
	u.Path = urlConsumeEvent

	request, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(jsonString))
	if err != nil {
		t.Error("NewRequest constructor error: ", err)
	}

	request.Header.Set("X-Custom-Header", "myvalue")
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		t.Error("NewRequest  error: ", err)
	}

	r, _ := readResponse(response)
	if r != "url1" {
		t.Error("error consume: ", r)
	}
}
