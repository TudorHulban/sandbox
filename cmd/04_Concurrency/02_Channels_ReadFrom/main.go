package main

import (
	"fmt"
	"sync"
)

var _poolWorkers int
var _mutex sync.Mutex

func doWork(work chan int, from, steps int) {
	for i := from; i <= from+steps; i++ {
		work <- i
	}

	_mutex.Lock()
	defer _mutex.Unlock()

	_poolWorkers--

	if _poolWorkers == 0 {
		fmt.Println("closing work")

		close(work)
	}
}

func main() {
	chWork := make(chan int)

	go doWork(chWork, 0, 10)
	_poolWorkers++

	go doWork(chWork, 11, 10)
	_poolWorkers++

	for {
		value, isOpen := <-chWork
		if !isOpen {
			break
		}

		fmt.Println("event:", value)
	}
}
