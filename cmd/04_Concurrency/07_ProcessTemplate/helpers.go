package main

import (
	"fmt"
	"sync"
	"time"
)

func dispatchWork(noWorkers int, s *service) {
	var wg sync.WaitGroup

	var wasPaused bool

	s.Start()

	for i := 0; i < noWorkers; i++ {
		wg.Add(1)

		if i > 3 && !wasPaused {
			wasPaused = true

			s.chPause <- struct{}{}
		}

		go s.DoWork(&wg)
	}

	fmt.Println("work created")

	go func() {
		time.Sleep(800 * time.Millisecond)
		s.chPause <- struct{}{}
	}()

	wg.Wait()

	s.chStop <- struct{}{}
	s.Stop()
	s.cleanUp()
}

func receiveWork(s *service) {
	var i int

loop:
	for {
		select {
		case work := <-s.chResults:
			fmt.Printf("ev%d: %d\n",
				i,
				work,
			)

			i++

		case <-s.chStop:
			s.Stop()

			break loop
		}
	}
}
