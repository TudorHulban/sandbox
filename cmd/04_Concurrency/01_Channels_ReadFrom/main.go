package main

import "fmt"

// we close the channel when go routine finishes
// and not when main go routine finishes.

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

		fmt.Printf("event: %d\n", value)
	}
}
