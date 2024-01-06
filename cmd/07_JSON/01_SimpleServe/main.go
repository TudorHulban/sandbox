package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("starting...")

	http.HandleFunc(mainRoute, handleRoutes)

	server := http.Server{
		Addr:              listeningOn,
		ReadHeaderTimeout: 3 * time.Second,
	}

	log.Fatal(
		server.ListenAndServe(),
	)
}
