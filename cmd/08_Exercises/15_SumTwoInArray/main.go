package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 2, 3}
	t := 5

	fmt.Println(indexes(a, t))
}

// indexes Returns the single pair matching target. No error handling.
func indexes(a []int, target int) [2]int {
	sort.Ints(a)

	for i := range a {
		for j := i + 1; j < len(a) && a[i]+a[j] <= target; j++ {
			if a[i]+a[j] == target {
				return [2]int{i, j}
			}
		}
	}

	return [2]int{}
}
