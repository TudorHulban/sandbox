package main

// file details filtering a slice of structs using a filtering function.

import (
	"log"
)

type person struct {
	name string
	age  int64
}

func main() {
	persons := []person{
		{name: "Baby", age: 4},
		{name: "Child", age: 8},
		{name: "GrownUp", age: 16},
	}

	log.Println("Persons:", persons)
	log.Println("Maximum 7", filterPersons(persons, filter7max))
	log.Println("Minimum 7", filterPersons(persons, filter7min))
}

func filterPersons(persons []person, f func(i int64) bool) []person {
	var result []person

	for _, person := range persons {
		if f(person.age) {
			result = append(result, person)
		}
	}

	return result
}

func filter7max(i int64) bool {
	return i < 7
}

func filter7min(i int64) bool {
	return i >= 7
}
