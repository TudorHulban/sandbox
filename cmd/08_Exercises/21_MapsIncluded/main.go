package main

import "fmt"

func main() {
	s1 := []string{"a", "d"}
	s2 := []string{"a", "b", "a"}

	fmt.Println(sliceIncluded(s1, s2))
}
