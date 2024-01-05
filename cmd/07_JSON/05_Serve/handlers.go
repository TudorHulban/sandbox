package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func serveError(w http.ResponseWriter, theError string) {
	serveJSON(w, "error -"+theError)
}

func serveJSON(w http.ResponseWriter, content string) {
	json2stream := json.NewEncoder(w)

	_ = json2stream.Encode(&content)
}

func collectionFetch(w http.ResponseWriter, r *http.Request) {
	fmt.Println("collectionFetch items:", _toServeItems)

	results, errMarshal := json.Marshal(_toServeItems)
	if errMarshal != nil {
		log.Println("json.Marshal: ", errMarshal)

		serveError(w, errMarshal.Error())

		return
	}

	fmt.Println("collectionFetch results: ", string(results))

	serveJSON(w, string(results))
}
