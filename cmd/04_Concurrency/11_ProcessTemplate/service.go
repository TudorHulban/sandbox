package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type ID int

type resource struct {
	id int64
}

type service struct {
	cache map[ID]*resource

	chResults chan int64
	chStop    chan struct{}
	chPause   chan struct{}

	mu sync.RWMutex
}

func newService() *service {
	return &service{
		cache: make(map[ID]*resource),

		chResults: make(chan int64),
		chStop:    make(chan struct{}),
		chPause:   make(chan struct{}),
	}
}

func (s *service) Start() error {
	return nil
}

func (s *service) DoWork(wg *sync.WaitGroup) {
	var isPaused bool

	toggle := func() {
		isPaused = !isPaused
	}

loop:
	for {
		select {
		case <-s.chPause:
			toggle()

			if isPaused {
				fmt.Println("service is now paused ....")
			} else {
				fmt.Println("service is now resumed ....")
			}

		default:
			if isPaused {
				continue
			}

			s.chResults <- s.createResource(ID(time.Now().UnixNano())).id

			wg.Done()

			break loop
		}
	}
}

func (s *service) GetStopCh() chan struct{} {
	return s.chStop
}

func (s *service) Stop() error {
	return nil
}

func (s *service) cleanUp() {
	close(s.chResults)
	close(s.chStop)
}

func (s *service) createResource(id ID) *resource {
	res := resource{
		id: time.Now().UnixNano(),
	}

	s.mu.Lock()
	s.cache[id] = &res
	s.mu.Unlock()

	return &res
}

func (s *service) getResource(id ID) (*resource, error) {
	s.mu.Lock()
	res, exists := s.cache[id]
	s.mu.Unlock()

	if !exists {
		return nil, fmt.Errorf("resource not found for ID: %d", id)
	}

	return res, nil
}

func (s *service) listResources() {
	s.mu.Lock()
	defer s.mu.Unlock()

	log.Println("number cached resources: ", len(s.cache))

	for id, resource := range s.cache {
		log.Printf(
			"resources with ID: %d and resource ID; %d",
			id,
			resource.id,
		)
	}
}

func (s *service) deleteResource(id ID) error {
	s.mu.Lock()
	_, exists := s.cache[id]
	s.mu.Unlock()

	if !exists {
		return fmt.Errorf("resource not found for ID: %d", id)
	}

	s.mu.Lock()
	delete(s.cache, id)
	s.mu.Unlock()

	return nil
}

func (s *service) cleanCache() {
	s.mu.Lock()
	s.cache = make(map[ID]*resource)
	s.mu.Unlock()
}
