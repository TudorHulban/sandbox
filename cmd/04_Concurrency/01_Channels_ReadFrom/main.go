package main

import (
	"log"
)

func main() {
	n := 10
	chIntegers := make(chan int)

	go func() {
		defer close(chIntegers)

		for i := 0; i < n; i++ {
			chIntegers <- i
		}
	}()

	for {
		value, isOpen := <-chIntegers
		if !isOpen {
			break
		}

		log.Printf("event: %d\n", value)
	}
}
