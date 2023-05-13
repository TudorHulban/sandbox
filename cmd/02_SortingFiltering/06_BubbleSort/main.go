// Bubble Sort in Golang
package main

import (
	"fmt"
)

func main() {
	a := []int{3, 2, 4, 1}
	fmt.Println("initial:", a)

	sort(a)
	fmt.Println("sorted:", a)
}

func sort(a []int) {
	isSorted := false

	for !isSorted {
		isSorted = true

		for i := 0; i < len(a)-1; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
				isSorted = false
			}
		}

		if isSorted {
			break
		}
	}
}
