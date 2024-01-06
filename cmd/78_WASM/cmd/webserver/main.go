package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	listeningOn := fmt.Sprintf(
		"localhost:%d",
		listeningPort,
	)

	fmt.Printf(
		"listening on http://%s/index.html\n",
		listeningOn,
	)

	server := http.Server{
		Addr:              listeningOn,
		ReadHeaderTimeout: 3 * time.Second,

		Handler: http.FileServer(
			http.Dir("../../assets"),
		),
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Failed to start server", err)

		return
	}
}
