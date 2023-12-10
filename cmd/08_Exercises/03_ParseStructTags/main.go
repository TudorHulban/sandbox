package main

import (
	"fmt"
)

type User struct {
	name string `hera:"fieldName1"`
	age  int    `hera:"fieldName2"`
}

func main() {
	fmt.Printf("Table Name: '%s'.\n", getTableName(&User{}))

	fields := getFields(&User{})

	for ix, field := range fields {
		fmt.Printf("Field %d: has field name: %v.\n", ix, field.FieldName)
	}
}
