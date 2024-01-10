package helpers

import "golang.org/x/exp/constraints"

func Median[T constraints.Integer](slice []T) float64 {
	if len(slice)%2 == 0 {
		middle := len(slice)/2 - 1

		return float64((slice[middle] + slice[middle+1])) / 2
	}

	return float64(slice[len(slice)/2])
}
