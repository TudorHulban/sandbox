package main

import (
	"log"
)

/*
Code gets number of pairs from an array.
A pair is a configurable number of occurrences.
*/

func main() {
	data := []int{1, 2, 1, 1, 3, 2, 4, 2, 2}

	log.Println(getPairs(2, data))
}
