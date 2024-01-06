package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func sortSliceOfArrays(arrays [][2]int) {
	var isSorted bool

	for !isSorted {
		isSorted = true

		for i := 0; i < len(arrays)-1; i++ {
			if arrays[i][0] > arrays[i+1][0] {
				arrays[i], arrays[i+1] = arrays[i+1], arrays[i]

				isSorted = false

				continue
			}
		}
	}
}

func TestSortSliceArrays(t *testing.T) {
	testCases := []struct {
		description string
		input       [][2]int
		result      [][2]int
	}{
		{
			"1. one",
			[][2]int{{1, 1}},
			[][2]int{{1, 1}},
		},
		{
			"2. two",
			[][2]int{{2, 1}, {1, 1}},
			[][2]int{{1, 1}, {2, 1}},
		},
		{
			"3. three",
			[][2]int{{-2, 6}, {4, 0}, {-2, 6}},
			[][2]int{{-2, 6}, {-2, 6}, {4, 0}},
		},
		{
			"4. six",
			[][2]int{{-2, 6}, {4, 0}, {4, 0}, {4, 0}, {4, 0}, {-2, 6}},
			[][2]int{{-2, 6}, {-2, 6}, {4, 0}, {4, 0}, {4, 0}, {4, 0}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description,
			func(t *testing.T) {
				sortSliceOfArrays(tc.input)

				assert.Equal(t, tc.result, tc.input)
			},
		)
	}
}
