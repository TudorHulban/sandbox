package main

import "fmt"

func main() {
	test := "The quick brown fox jumps over the lazy dog."

	fmt.Printf(
		"number words for: %s\nis %d.\n",
		test,
		noWords(test),
	)
}
