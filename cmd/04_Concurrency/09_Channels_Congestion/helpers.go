package main

import (
	"log"
	"strconv"
)

func createWork(noTasks, maxConcurrentTasks uint, tasks chan task) {
	if noTasks == 0 {
		return
	}

	var currentLoad uint

	for noTasks > 0 {
		currentLoad = uint(len(tasks))

		log.Println(
			"creating work, remaining tasks:", noTasks,
			"maxNoConcTasks:", maxConcurrentTasks,
			"load:", currentLoad)

		if currentLoad > maxConcurrentTasks {
			log.Printf("current load('%d') exceeding max number of conncurrent tasks: %d",
				currentLoad,
				maxConcurrentTasks,
			)

			continue // allowing time for tasks to finish
		}

		log.Printf(
			"channel load: %d",
			currentLoad,
		)

		request := "R" + strconv.Itoa(int(noTasks))

		tasks <- task{request: request}

		log.Println("request: ", request)

		noTasks--
	}

	log.Printf("no more work, remaining tasks: %d", noTasks)
}
