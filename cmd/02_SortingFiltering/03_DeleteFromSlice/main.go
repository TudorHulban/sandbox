package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2}
	fmt.Printf("array source: %v\n", a)

	b, _ := delIndex(a, 1)
	fmt.Printf("array destination: %v\n", b)
}

func delIndex(slice []int, index int) ([]int, error) {
	if index > len(slice) {
		return nil, fmt.Errorf("maximum index is %d, passed value is %d", len(slice), index)
	}

	return append(slice[:index], slice[index+1:]...), nil
}

func delIndexG[T any](slice []T, index int) ([]T, error) {
	if index > len(slice) {
		return nil, fmt.Errorf("maximum index is %d, passed value is %d", len(slice), index)
	}

	return append(slice[:index], slice[index+1:]...), nil
}
