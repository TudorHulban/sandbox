package main

import (
	"log"
	"reflect"
	"strconv"
)

type Task struct {
	ID   int
	Name string
	Area string
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
	for k, v := range w.Tasks {
		if v.ID == id {
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
	var res []interface{}

	for _, v := range w.Tasks {
		res = append(res, reflect.ValueOf(v).FieldByName(fieldName).Interface())

	}

	return res
}

func showFieldDirect(fieldName string, work *Work) []interface{} {
	var res []interface{}

	for _, v := range work.Tasks {
		res = append(res, reflect.ValueOf(v).FieldByName(fieldName).Interface())
	}

	return res
}

func showFieldName(fieldName string, work *Work) []interface{} {
	var res []interface{}

	for _, v := range work.Tasks {
		res = append(res, v.Name)
	}

	return res
}

func showFieldAssertion(fieldName string, from interface{}) []interface{} {
	var res []interface{}

	for _, v := range from.(Work).Tasks {
		res = append(res, reflect.ValueOf(v).FieldByName(fieldName).Interface())
	}

	return res
}
