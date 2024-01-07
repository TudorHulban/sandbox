package main

import (
	"fmt"
)

// var _operationNumber int

func main() {
	h := newHanoi()
	fmt.Println(h)
}

// func moveDisks(numDisks int, from, to, via *rod) {
// 	if numDisks == 0 || len(*from) == 0 {
// 		return
// 	}

// 	if len(*from) == 1 {
// 		moveDisk(from, to)

// 		return
// 	}

// 	moveDisks(numDisks-1, from, to, via)
// 	moveDisk(from, to)
// }

// func moveDisk(from, to *rod) {
// 	if len(*from) == 0 {
// 		return
// 	}

// 	_operationNumber++
// 	fmt.Println(_operationNumber, "Before:", "rodFrom: ", rodFrom, "rodAux: ", rodAux, "rodTo: ", rodTo)

// 	*to = append(*to, (*from)[len(*from)-1])
// 	*from = (*from)[:len(*from)-1]

// 	fmt.Println(_operationNumber, "After :", "rodFrom: ", rodFrom, "rodAux: ", rodAux, "rodTo: ", rodTo)
// }
