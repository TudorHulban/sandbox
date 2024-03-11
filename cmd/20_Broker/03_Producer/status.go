package main

import (
	"log"
	"time"
)

func getStatus(forRequest request) {
	wait := time.NewTimer(20 * time.Second)
	<-wait.C

	log.Println("ready request: ", forRequest.id, readyRequest(forRequest))
}
