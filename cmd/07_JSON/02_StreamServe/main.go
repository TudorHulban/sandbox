package main

import (
	"fmt"

	"net/http"
	"time"
)

type httpResponse struct {
	Msg string `json:"message"`
}

func main() {
	http.HandleFunc("/x", handlerOne)
	http.HandleFunc("/y", handlerTwo)

	server := &http.Server{
		Addr:              fmt.Sprintf(":%v", port),
		ReadHeaderTimeout: 3 * time.Second,
	}

	_ = server.ListenAndServe()
}
