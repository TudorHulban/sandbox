package main

import (
	"log"
	"strconv"
	"time"
)

type task struct {
	request string
}

type pool struct {
	chInputs  chan task
	chOutputs chan string
	noWorkers int
	noTasks   int
}

func (p *pool) do(workerID int) {
	for j := range p.chInputs {
		log.Printf("worker %d started job %s.\n", workerID, j)

		// long process for work
		time.Sleep(1 * time.Millisecond)

		log.Printf("worker %d finished job %s.\n", workerID, j)

		p.chOutputs <- "finished: " + strconv.FormatInt(time.Now().UnixNano(), 10) + " " + j.request

		log.Printf("work sent by worker %d for request %s.\n", workerID, j.request)
	}

	log.Printf("closing worker %d.\n", workerID)
}

func (p *pool) cleanUp() {
	close(p.chOutputs)
	close(p.chInputs)
}
