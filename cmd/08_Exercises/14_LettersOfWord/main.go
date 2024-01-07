package main

import (
	"fmt"
	"sort"
)

func main() {
	s := "abcd"
	letters := make(map[letter]bool)

	deleteOne(s, &letters)

	order := make([][]string, len(s)+1)

	for element := range letters {
		order[len(element)] = append(order[len(element)],
			string(element),
		)
	}

	order[len(s)] = append(order[len(s)], s)

	for ix := range order {
		if ix == 0 {
			continue
		}

		sort.Strings(order[ix])

		fmt.Println(order[ix])
	}
}
