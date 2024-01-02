package main

import (
	"log"
	"reflect"
)

func main() {
	t1 := Task{ID: 1, Name: "x1", Area: "db"}
	t2 := Task{ID: 2, Name: "x2", Area: "os"}
	t3 := Task{ID: 3, Name: "x3", Area: "db"}

	var w1 Work
	w1.Tasks = []Task{t1, t2, t3}

	log.Println(w1.contains(11))
	log.Println(w1.contains(1))

	w2 := *NewWork(10)
	log.Println(w2.contains(11))
	log.Println(w2.contains(1))

	log.Println(w2.showFieldValues("Name"))
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
