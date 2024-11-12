package main

import (
	"fmt"
	"time"
)

func worker(chPayload chan int) {
	for _ = range chPayload {
		// fmt.Println(work)
	}
}

func dispatcher(noWorkers, load int) {
	now := time.Now()

	chPayload := make(chan int)

	for range noWorkers {
		go worker(chPayload)
	}

	for payload := range load {
		chPayload <- payload
	}

	close(chPayload)

	fmt.Printf(
		"no workers: %d - microseconds: %d\n",

		noWorkers,
		time.Since(now).Microseconds(),
	)
}

func main() {
	noWorkers := 3
	load := 10

	dispatcher(noWorkers, load)
}
