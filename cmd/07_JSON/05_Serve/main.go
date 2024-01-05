package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	_toServeItems.Items = append(_toServeItems.Items,
		item{"new", "x"},
	)
	_toServeItems.Items = append(_toServeItems.Items,
		item{"new", "y"},
	)

	fmt.Println("items:", _toServeItems)

	r := mux.NewRouter()

	r.HandleFunc("/"+theItem, collectionFetch).Methods("GET")
	http.Handle("/", r)

	srv := http.Server{
		Handler:      r,
		Addr:         theURL,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("serving on %s/tasks\n", theURL)

	log.Fatal(srv.ListenAndServe())
}
