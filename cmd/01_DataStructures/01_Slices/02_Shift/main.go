package main

import "fmt"

func shiftRight(slice *[]int) {
	*slice = append((*slice)[len(*slice)-1:], (*slice)[:len(*slice)-1]...)
}

func main() {
	s := []int{5, 7, 8, 4, 9, 1}

	numberTimes := 2

	for i := 0; i < numberTimes; i++ {
		shiftRight(&s)
	}

	fmt.Println(s) // [9 1 5 7 8 4]
}
