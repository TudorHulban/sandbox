package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type task struct {
	begin int
	end   int
}

type tasks []task

type pool struct {
	chWork chan int
	mu     sync.Mutex

	noWorkers int
	noTasks   int

	tasksHandled []int
}

func (p *pool) dispatchWork(from, steps int) {
	tasks := createWork(from, steps, p.noWorkers)

	for _, task := range tasks {
		go p.do(task)

		p.noTasks++
	}
}

func (p *pool) do(t task) {
	for i := t.begin; i <= t.end; i++ {
		time.Sleep(1 * time.Millisecond) // simulate some work

		p.chWork <- i
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	p.noTasks--

	if p.noTasks == 0 {
		log.Println("no more tasks. closing work.")

		close(p.chWork)

		log.Println("work now closed.")
	}
}

func (p *pool) start() {
	for {
		event, isOpen := <-p.chWork
		if !isOpen {
			fmt.Println("work is not opened anymore, exiting...")

			break
		}

		p.tasksHandled = append(p.tasksHandled, event)

		log.Printf(
			"done work %v.",
			event,
		)
	}
}
