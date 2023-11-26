package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("listening on http://localhost:9090/index.html")

	err := http.ListenAndServe(
		":9090",
		http.FileServer(http.Dir("../../assets")),
	)
	if err != nil {
		fmt.Println("Failed to start server", err)

		return
	}
}
