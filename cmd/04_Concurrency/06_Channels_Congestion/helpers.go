package main

import (
	"log"
	"strconv"
)

func createWork(noTasks int, tasks chan task, maxNoConcTasks int) {
	if noTasks < 1 {
		return
	}

	var load int

	for noTasks > 0 {
		load = len(tasks)

		log.Println("creating work, remaining tasks: ", noTasks, "load:", load)

		if load <= maxNoConcTasks {
			log.Printf("channel load: %d", load)

			request := "R" + strconv.Itoa(noTasks)

			tasks <- task{request: request}

			log.Println("request: ", request)

			noTasks--
		}
	}

	log.Printf("no more work, remaining tasks: %d", noTasks)
}
