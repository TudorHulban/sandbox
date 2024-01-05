package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

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
