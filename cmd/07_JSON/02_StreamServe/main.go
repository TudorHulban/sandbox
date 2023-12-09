package main

import (
	"encoding/json"
	"fmt"

	"log"
	"net/http"
	"time"
)

type httpResponse struct {
	Msg string `json:"message"`
}

func main() {
	http.HandleFunc("/x", handlerOne)
	http.HandleFunc("/y", handlerTwo)

	_ = http.ListenAndServe(
		fmt.Sprintf(":%v", port),
		nil,
	)
}

func handlerOne(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	response := httpResponse{
		Msg: msgOne,
	}

	marshalledResponse, errMarshal := json.Marshal(response)
	if errMarshal != nil {
		log.Println(errMarshal)
	}

	fmt.Fprint(w, string(marshalledResponse))

	fmt.Println("duration:", time.Since(startTime))
}

func handlerTwo(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	response := httpResponse{
		Msg: msgTwo,
	}

	json2stream := json.NewEncoder(w)
	_ = json2stream.Encode(&response)

	fmt.Println("duration:", time.Since(startTime))
}
