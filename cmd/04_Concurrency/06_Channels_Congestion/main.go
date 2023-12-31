package main

import (
	"log"
	"math"
	"time"
)

func main() {
	p := pool{
		noWorkers: 3,
		noTasks:   10,
		chOutputs: make(chan string),
	}

	// could be different than number of workers
	maxConcurrentTasks := uint(math.Min(
		float64(p.noTasks),
		3.0,
	))
	p.chInputs = make(chan task, maxConcurrentTasks)

	defer p.cleanUp()

	timeStart := time.Now()

	for workerID := 0; workerID < p.noWorkers; workerID++ {
		go p.do(workerID)

		log.Println("created worker:", workerID)
	}

	go createWork(
		p.noTasks,
		maxConcurrentTasks,
		p.chInputs,
	)

	result := []string{"\n"}

	for r := 0; r < int(p.noTasks); r++ {
		result = append(
			result,
			<-p.chOutputs+"\n",
		)
	}

	log.Println(result)
	log.Println("elapsed: ", time.Since(timeStart))
}
