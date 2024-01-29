package ddltable

import "reflect"

func getObjectName(object any) string {
	t := reflect.TypeOf(object)

	var result string

	for t.Kind() == reflect.Ptr {
		t = t.Elem()

		result = result + "*"
	}

	return result + t.Name()
}
