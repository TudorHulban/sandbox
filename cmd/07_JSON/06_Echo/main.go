package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type httpMessage struct {
	Text string `json:"message"`
}

func main() {
	http.HandleFunc("/url1", responseText)
	http.HandleFunc("/url2", responseJSON)
	http.HandleFunc("/url3", responseString)

	log.Println("Listening:", port)

	server := &http.Server{
		Addr:              ":" + port,
		ReadHeaderTimeout: 3 * time.Second,
	}

	log.Fatal(
		server.ListenAndServe(),
	)
}

// handleURL test with:
// curl -X POST -H 'Content-Type:
// application/json' -d "{\"message\": \"Hi\"}" http://127.0.0.1:8080/url1
func responseText(w http.ResponseWriter, r *http.Request) {
	log.Println("new request ...")

	decoder := json.NewDecoder(r.Body)

	var event httpMessage

	errDecodeEvent := decoder.Decode(&event)
	if errDecodeEvent != nil {
		log.Fatal("decode event: ", errDecodeEvent)

		return
	}

	log.Println("response text: ", event.Text)

	_, _ = w.Write(
		[]byte(event.Text),
	)
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
