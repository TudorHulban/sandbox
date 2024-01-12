package main

import (
	"fmt"
)

func main() {
	number := 2

	fmt.Println(integerToBinary(number))
	fmt.Println(integerToBinary(number - 1))

	fmt.Println(
		isPowerOfTwo(number),
	)
}
