package main

import (
	"fmt"
)

type User struct {
	Name string `hera:"fieldName1"`
	Age  int    `hera:"fieldName2"`
}

func main() {
	fmt.Printf("Table Name: '%s'.\n", getTableName(&User{}))

	fields := getFields(&User{})

	for ix, field := range fields {
		fmt.Printf(
			"field %d: has field name: %v.\n",
			ix,
			field.FieldName,
		)
	}
}
