package main

import (
	"fmt"
)

func main() {
	words := []string{"cat", "baby", "dog", "bird", "car"}
	scrambled1 := "bbabylkkjabdog"
	scrambled2 := "xxx"

	fmt.Println("result:", parse(scrambled1, words...))
	fmt.Println("result:", parse(scrambled2, words...))
}
