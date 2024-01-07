package main

import (
	"fmt"
)

func main() {
	fmt.Println(factorialImperative(3))
	fmt.Println(factorialRecursive(3))
}

func factorialImperative(n int) int {
	result := 1

	for i := 1; i <= n; i++ {
		result = result * i
	}

	return result
}

func factorialRecursive(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorialRecursive(n-1)
}
