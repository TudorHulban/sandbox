package main

import "fmt"

var _operationNumber int

func moveDisk(from, to *rod) {
	if len(*from) == 0 {
		return
	}

	_operationNumber++

	fmt.Printf("%d. before: rodFrom %d - rodTo %d ",
		_operationNumber,
		*from,
		*to,
	)

	*to = append(*to, (*from)[len(*from)-1])
	*from = (*from)[:len(*from)-1]

	fmt.Printf("- after: rodFrom %d - rodTo %d\n", *from, *to)
}

func moveDisks(numDisks int, from, to, via *rod) {
	if numDisks == 0 || len(*from) == 0 {
		return
	}

	if len(*from) == 1 {
		moveDisk(from, to)

		return
	}

	moveDisks(numDisks-1, from, to, via)
	moveDisk(from, via)
}
