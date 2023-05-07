package main

import (
	"log"
	"time"
)

func main() {
	p := pool{
		noWorkers: 5,
		chWork:    make(chan int),
	}

	timeStart := time.Now()

	p.dispatchWork(0, 25)
	p.start()

	timeEnd := time.Now()
	log.Printf("Duration: %s.", timeEnd.Sub(timeStart))
}
