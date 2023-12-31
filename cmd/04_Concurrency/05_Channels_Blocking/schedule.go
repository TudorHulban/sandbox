package main

import (
	"log"
	"time"
)

func schedule(what func(), interval time.Duration) chan struct{} {
	chStop := make(chan struct{})

	go func() {
		for {
			what()

			select {
			case <-time.After(interval):
				log.Println("interval finished, triggering event")

			case <-chStop:
				log.Println("exiting scheduler goroutine")

				return
			}
		}
	}()

	return chStop
}

func cleanup() {
	log.Println("cleanup")

	time.Sleep(2 * time.Second) // delay exiting for catching previous sends on channel and logs.
}
