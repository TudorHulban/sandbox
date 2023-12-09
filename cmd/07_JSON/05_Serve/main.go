package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	_toServeItems.Items = append(_toServeItems.Items, item{"new", "x"})
	_toServeItems.Items = append(_toServeItems.Items, item{"new", "y"})

	fmt.Println("items:", _toServeItems)

	r := mux.NewRouter()
	r.HandleFunc("/"+theItem, collectionFetch).Methods("GET")
	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         theURL,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func collectionFetch(w http.ResponseWriter, r *http.Request) {
	fmt.Println("collectionFetch items:", _toServeItems)

	results, errMarshal := json.Marshal(_toServeItems)
	if errMarshal != nil {
		log.Println("json.Marshal: ", errMarshal)

		serveError(w, errMarshal.Error())

		return
	}

	fmt.Println("collectionFetch results: ", string(results[:]))

	serveJSON(w, string(results))
}

func serveError(w http.ResponseWriter, theError string) {
	serveJSON(w, "error -"+theError)
}

func serveJSON(w http.ResponseWriter, content string) {
	json2stream := json.NewEncoder(w)

	json2stream.Encode(&content)
}
