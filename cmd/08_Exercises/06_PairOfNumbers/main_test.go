package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPairs(t *testing.T) {
	testCases := []struct {
		description string
		slice1      []int
		slice2      []int
		amount      int

		result [][2]int
	}{
		{
			"1. minimal happy path",
			[]int{1},
			[]int{1},
			2,
			[][2]int{{1, 1}},
		},
		{
			"2. no result",
			[]int{1},
			[]int{0},
			2,
			nil,
		},
		{
			"3. happy path",
			[]int{-1, -2, 4, 4, -2, -6, 5, -7},
			[]int{0, 6, 3, 4, 0},
			4,
			[][2]int{{-2, 6}, {-2, 6}, {4, 0}, {4, 0}, {4, 0}, {4, 0}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description,
			func(t *testing.T) {
				array := sumOfTwo(
					tc.slice1,
					tc.slice2,
					tc.amount,
				)

				t.Log("unsorted result:", array)

				sortSliceOfArrays(array)

				assert.Equal(t,
					tc.result,
					array,

					"sumOfTwo",
				)
			},
		)
	}
}
