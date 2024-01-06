package main

import "sort"

func orderWStandardLibSort(a1, a2 []int) []int {
	result := append(a1, a2...)

	sort.Ints(result)

	return result
}

func orderWComparison(slice1, slice2 []int) []int {
	var result []int

	var indexSlice2 int

	for _, valueSlice1 := range slice1 {
		for indexSlice2 <= len(slice2)-1 {
			if valueSlice1 < slice2[indexSlice2] {
				break
			}

			result = append(result,
				slice2[indexSlice2],
			)

			indexSlice2++
		}

		result = append(result,
			valueSlice1,
		)
	}

	return append(result,
		slice2[indexSlice2:]...,
	)
}
