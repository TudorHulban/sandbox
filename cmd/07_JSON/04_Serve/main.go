package main

import (
	"log"
)

func main() {
	log.Println(
		"DoPOST: ",
		DoPOST(
			"http://localhost:3002",
			NewConfig(4),
		),
	)
}
