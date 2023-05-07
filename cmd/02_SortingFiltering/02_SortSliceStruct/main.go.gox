package main

import (
	"fmt"
	"sort"
)

type task struct {
	id   int
	name string
}

// no pointer needed due to slice mechanics
func sortAsc(tasks []task) {
	asc := func(i, j int) bool {
		return tasks[i].id < tasks[j].id
	}

	sort.Slice(tasks, asc)
}

func sortName(tasks []task) {
	byName := func(i, j int) bool {
		return tasks[i].name < tasks[j].name
	}

	sort.Slice(tasks, byName)
}

func main() {
	tasks := []task{
		{id: 11, name: "t11"},
		{id: 2, name: "t2"},
		{id: 3, name: "t4"},
		{id: 3, name: "t3"},
		{id: 4, name: "t1"},
	}
	fmt.Printf("Raw: %v\n", tasks)

	sortAsc(tasks)
	fmt.Printf("Sort ascending: %v\n", tasks)

	sortName(tasks)
	fmt.Printf("Sort by name: %v\n", tasks)
}
