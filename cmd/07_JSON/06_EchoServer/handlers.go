package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// handleURL test with:
// curl -X POST -H 'Content-Type:
// application/json' -d "{\"message\": \"Hi\"}" http://127.0.0.1:8080/url1
func responseText(w http.ResponseWriter, r *http.Request) {
	log.Println("new request ...")

	decoder := json.NewDecoder(r.Body)

	var event httpMessage

	if errDecodeEvent := decoder.Decode(&event); errDecodeEvent != nil {
		log.Fatal("decode event: ", errDecodeEvent)

		return
	}

	log.Println("response text: ", event.Text)

	if _, errWrite := w.Write(
		[]byte(event.Text),
	); errWrite != nil {
		fmt.Println("errWrite:", errWrite)
	}
}

// handleURL test with:
// curl -X POST -H 'Content-Type:
// application/json' -d "{\"message\": \"Hi\"}" http://127.0.0.1:8080/url2
func responseJSON(w http.ResponseWriter, r *http.Request) {
	log.Println("new request ...")

	decoder := json.NewDecoder(r.Body)

	var event httpMessage

	errDecode := decoder.Decode(&event)
	if errDecode != nil {
		log.Fatal("decode event: ", errDecode)

		return
	}

	log.Println("response JSON: ", event.Text)

	rawResponse := httpMessage{
		Text: event.Text,
	}

	json2stream := json.NewEncoder(w)
	_ = json2stream.Encode(&rawResponse)
}

// handleURL test with:
// curl -X POST -H 'Content-Type:
// application/json' -d "{\"message\": \"Hi\"}" http://127.0.0.1:8080/url3
func responseString(w http.ResponseWriter, r *http.Request) {
	log.Println("new request ...")

	decoder := json.NewDecoder(r.Body)

	var event httpMessage

	errDecode := decoder.Decode(&event)
	if errDecode != nil {
		log.Fatal("error decode: ", errDecode)

		return
	}

	log.Println("response: ", event.Text)

	var shellResponse bytes.Buffer

	rawResponse := httpMessage{
		Text: event.Text,
	}

	json2stream := json.NewEncoder(&shellResponse)
	_ = json2stream.Encode(&rawResponse)

	_, _ = w.Write(
		shellResponse.Bytes(),
	)
}
