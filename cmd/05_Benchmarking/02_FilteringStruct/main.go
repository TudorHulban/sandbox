package main

import (
	"log"
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

func main() {
	t1 := Task{ID: 1, Name: "x1", Area: "db"}
	t2 := Task{ID: 2, Name: "x2", Area: "os"}
	t3 := Task{ID: 3, Name: "x3", Area: "db"}

	var w1 Work
	w1.Tasks = []Task{t1, t2, t3}

	log.Println(w1.contains(11))
	log.Println(w1.contains(1))

	w2 := *NewTasks(10)
	log.Println(w2.contains(11))
	log.Println(w2.contains(1))

	log.Println(w2.showFieldValues("Name"))
}

func (w *Work) contains(id int) int {
	for k, task := range w.Tasks {
		if task.ID == id {
			return k
		}
	}

	return 0
}

func NewTasks(howMany int) *Work {
	var work Work

	for i := 0; i < howMany; i++ {
		task := Task{
			ID:   i,
			Name: "N" + strconv.Itoa(i),
			Area: "A" + strconv.Itoa(i),
		}

		work.Tasks = append(work.Tasks, task)
	}

	return &work
}

func (w *Work) showFieldValues(fieldName string) []interface{} {
	result := make([]interface{}, len(w.Tasks))

	for ix, v := range w.Tasks {
		result[ix] = reflect.ValueOf(v).FieldByName(fieldName).Interface()
	}

	return result
}

func showFieldDirect(fieldName string, work *Work) []interface{} {
	result := make([]interface{}, len(work.Tasks))

	for ix, task := range work.Tasks {
		result[ix] = reflect.ValueOf(task).FieldByName(fieldName).Interface()
	}

	return result
}

func showFieldName(fieldName string, work *Work) []interface{} {
	result := make([]interface{}, len(work.Tasks))

	for ix, task := range work.Tasks {
		result[ix] = task.Name
	}

	return result
}

func showFieldAssertion(fieldName string, from any) []any {
	var res []any //nolint:prealloc

	for _, task := range from.(*Work).Tasks {
		res = append(res,
			reflect.ValueOf(task).FieldByName(fieldName).Interface(),
		)
	}

	return res
}
