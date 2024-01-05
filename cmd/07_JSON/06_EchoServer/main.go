package main

import (
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
