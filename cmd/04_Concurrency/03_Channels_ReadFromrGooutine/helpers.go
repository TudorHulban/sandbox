package main

import (
	"log"
	"math"
)

func createWork(from, steps, noWorkers int) tasks {
	load := int(math.Floor(float64(steps) / float64(noWorkers)))

	var lastStep int
	var assigned int
	var res tasks

	for i := 0; i < noWorkers; i++ {
		var t task

		switch i {
		case 0:
			{
				t.begin = 0
				lastStep = load
			}

		default:
			{
				t.begin = lastStep + 1
				lastStep = lastStep + load
			}
		}

		assigned = assigned + load

		t.end = lastStep
		res = append(res, t)
	}

	res[noWorkers-1].end = res[noWorkers-1].end + (steps - assigned)

	for ix, task := range res {
		log.Printf("Task %v with following details: begin - %v, end - %v.\n",
			ix+1, task.begin, task.end)
	}

	return res
}
