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

	noBuffer := int(math.Min(float64(p.noTasks), 10.0))
	p.chInputs = make(chan task, noBuffer)

	defer p.cleanUp()

	timeStart := time.Now()

	for w := 0; w < p.noWorkers; w++ {
		go p.do(w)

		log.Println("created worker:", w)
	}

	go createWork(p.noTasks, p.chInputs, noBuffer)

	res := []string{"\n"}

	for r := 0; r < p.noTasks; r++ {
		res = append(res, <-p.chOutputs+"\n")
	}

	log.Println(res)
	log.Println("elapsed: ", time.Since(timeStart))
}
