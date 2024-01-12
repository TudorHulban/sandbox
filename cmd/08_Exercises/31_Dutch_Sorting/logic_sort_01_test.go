package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort01(t *testing.T) {
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
			"2. 10",
			[]int{1, 0},
			[]int{0, 1},
		},
		{
			"3. 1111",
			[]int{1, 1, 1, 1},
			[]int{1, 1, 1, 1},
		},
		{
			"4. 010",
			[]int{0, 1, 0},
			[]int{0, 0, 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description,
			func(t *testing.T) {
				assert.Equal(t,
					tc.want,
					sort01(tc.input))
			},
		)
	}
}
