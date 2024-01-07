package main

import (
	"fmt"
)

func main() {
	h := newHanoi()

	fmt.Println(h)

	moveDisks(
		h.numDisks,
		&h.rodFrom,
		&h.rodTo,
		&h.rodAux,
	)
}
