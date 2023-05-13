package main

import "fmt"

// data Structure to contain priority data that is related to heap.
type data struct {
	priority int
}

// Priority Structure to keep heap.
type queue struct {
	state []*data
}

// NewPriority Constructor for priority queue kind of type.
func NewQueue() *queue {
	return &queue{
		state: []*data{},
	}
}

// Insert Method adds data to the queue.
func (q *queue) Insert(d *data) *queue {
	q.state = append(q.state, d)

	if len(q.state) == 1 {
		return q
	}

	q.heapifyMax(len(q.state) - 1)

	go fmt.Println(q.Values())

	return q
}

// Extract Method extracts top priority data from queue.
func (q *queue) Extract() *data {
	l := len(q.state)

	if l == 0 {
		return nil
	}

	if l == 1 {
		res := (*q).state[0]

		q.state = []*data{}
		return res
	}

	res := (*q).state[0]
	go fmt.Println("extracted:", res.priority)

	// swap root with last element
	q.state[0] = q.state[l-1]
	// chop array
	q.state = q.state[:l-1]

	q.heapifyDown(0)
	go fmt.Println(q.Values())

	return res
}

// heapifyDown Helper which roots element with top priority.
func (q *queue) heapifyDown(index int) {
	lastIndex := len(q.state) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = 1
		} else if q.state[l].priority > q.state[r].priority {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if q.state[index].priority < q.state[childToCompare].priority {
			q.swap(index, childToCompare)

			index = childToCompare
			l, r = left(index), right(index)
			continue
		}

		return
	}
}

// Values Method can be used to nicely print priority values.
func (q queue) Values() []int {
	res := make([]int, len(q.state))

	for i, data := range q.state {
		res[i] = data.priority
	}

	return res
}

func (q *queue) heapifyMax(index int) {
	for q.state[parent(index)].priority < q.state[index].priority {
		q.swap(parent(index), index)

		index = parent(index)
	}
}

func (q *queue) swap(i1, i2 int) {
	q.state[i1], q.state[i2] = q.state[i2], q.state[i1]
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}
