package main

// Quick Sort in Golang

import (
	"log"
)

func main() {
	slice := generateSlice(10)

	log.Println("Unsorted slice: ", slice)
	quicksort(slice)
	log.Println("Sorted slice: ", slice)
}
