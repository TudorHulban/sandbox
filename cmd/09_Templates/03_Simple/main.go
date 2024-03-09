package main

import (
	"os"
)

func main() {
	person := Person{
		Name:  "John",
		Age:   16,
		Tasks: []string{"T1", "T2", "T3"},
	}

	w := os.Stdout

	f1(w, "template_name", person)
	f2(w, "t2.tmpl", person)
	f3(w, "t2.tmpl", "t2.csv", person)
}
