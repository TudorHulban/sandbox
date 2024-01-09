package main

import "slices"

func parseBrute[T uint | int](arrays ...[]T) []T {
	var ix int

	result := make([]T, 0)

	for {
		var addedElement bool

		for _, slice := range arrays {
			if ix > len(slice)-1 {
				continue
			}

			result = append(result,
				(slice[ix]),
			)

			addedElement = true
		}

		if !addedElement {
			break
		}

		ix++
	}

	slices.Sort(result)

	return result
}
