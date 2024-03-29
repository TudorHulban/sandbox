package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		description string
		input       []int
		sumAmount   int
		want        [2]int
	}{
		{
			"1. fail fast",
			[]int{3, 3},
			7,

			[2]int{},
		},
		{
			"2. fail fast",
			[]int{3},
			6,

			[2]int{},
		},
		{
			"3. case",
			[]int{2, 7, 11, 15},
			9,

			[2]int{0, 1},
		},
		{
			"4. case",
			[]int{3, 2, 4},
			6,

			[2]int{1, 2},
		},
		{
			"5. case",
			[]int{3, 3},
			6,

			[2]int{0, 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description,
			func(t *testing.T) {
				assert.Equal(t,
					tc.want,
					findValues(tc.input, tc.sumAmount),
				)
			},
		)
	}
}
