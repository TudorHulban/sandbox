package main

import (
	"reflect"
	"strconv"
)

type Task struct {
	Name string
	Area string

	ID int
}

type Work struct {
	Tasks []Task
}

func NewWork(howManyTasks int) *Work {
	var work Work

	for i := 0; i < howManyTasks; i++ {
		task := Task{
			ID:   i,
			Name: "N" + strconv.Itoa(i),
			Area: "A" + strconv.Itoa(i),
		}

		work.Tasks = append(work.Tasks, task)
	}

	return &work
}

func (w *Work) contains(taskID int) int {
	for ix, task := range w.Tasks {
		if task.ID == taskID {
			return ix
		}
	}

	return 0
}

func (w *Work) showFieldValues(fieldName string) any {
	result := make([]any, len(w.Tasks))

	for ix, task := range w.Tasks {
		result[ix] = reflect.
			ValueOf(task).
			FieldByName(fieldName).
			Interface()
	}

	return result
}
