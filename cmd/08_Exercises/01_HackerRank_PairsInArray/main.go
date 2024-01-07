package main

import (
	"log"
)

func main() {
	log.Println(
		getPairs(
			2,
			[]int{1, 2, 1, 1, 3, 2, 4, 2, 2},
		),
	)
}
