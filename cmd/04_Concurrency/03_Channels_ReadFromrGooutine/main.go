package main

import (
	"fmt"
	"time"
)

func main() {
	p := pool{
		noWorkers: 5,
		chWork:    make(chan int),
	}

	timeStart := time.Now()

	p.dispatchWork(0, 24)
	p.start()

	fmt.Printf(
		"Duration: %s.\nNumber of tasks handled: %d.\n",
		time.Since(timeStart),
		len(p.tasksHandled),
	)
}
