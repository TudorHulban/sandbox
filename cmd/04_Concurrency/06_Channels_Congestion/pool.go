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
	noTasks   uint
}

func (p *pool) do(workerID int) {
	for job := range p.chInputs {
		log.Printf(
			"worker %d started job %s.\n",
			workerID,
			job,
		)

		time.Sleep(1 * time.Millisecond) // long process for work

		log.Printf(
			"worker %d finished job %s.\n",
			workerID,
			job,
		)

		p.chOutputs <- "finished: " +
			strconv.FormatInt(time.Now().UnixNano(), 10) +
			" " +
			job.request

		log.Printf(
			"work sent by worker %d for request %s.\n",
			workerID,
			job.request,
		)
	}

	log.Printf(
		"closing worker %d.\n",
		workerID,
	)
}

func (p *pool) cleanUp() {
	close(p.chOutputs)
	close(p.chInputs)
}
