package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort012(t *testing.T) {
	testCases := []struct {
		description string
		input       []int
		want        []int
	}{
		{
			"1. 0",
			[]int{0},
			[]int{0},
		},
		{
			"2. 0000",
			[]int{0, 0, 0, 0},
			[]int{0, 0, 0, 0},
		},
		{
			"3. 1111",
			[]int{1, 1, 1, 1},
			[]int{1, 1, 1, 1},
		},
		{
			"4. 2222",
			[]int{2, 2, 2, 2},
			[]int{2, 2, 2, 2},
		},
		{
			"5. 021",
			[]int{0, 2, 1},
			[]int{0, 1, 2},
		},
		{
			"6. 021001221",
			[]int{0, 2, 1, 0, 0, 1, 2, 2, 1},
			[]int{0, 0, 0, 1, 1, 1, 2, 2, 2},
		},
		{
			"7. 20110022",
			[]int{2, 0, 1, 1, 0, 0, 2, 2},
			[]int{0, 0, 0, 1, 1, 2, 2, 2},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description,
			func(t *testing.T) {
				assert.Equal(t,
					tc.want,
					sort012(tc.input),
				)
			},
		)
	}
}
