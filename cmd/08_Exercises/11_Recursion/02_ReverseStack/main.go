package main

// race conditions are ignored.

import (
	"fmt"
)

func main() {
	var s stack

	s.push(1)
	s.push(2)
	s.push(3)
	s.push(4)

	fmt.Println(s)

	s.reverse()

	fmt.Println(s)
}
