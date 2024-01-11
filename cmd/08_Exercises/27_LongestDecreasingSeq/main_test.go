package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestDecreaseValue(t *testing.T) {
	testCases := []struct {
		description string
		input       []int
		want        int
	}{
		{
			"1. empty",
			[]int{},
			0,
		},
		{
			"2. short",
			[]int{1},
			1,
		},
		{
			"3. two",
			[]int{1, 2},
			1,
		},
		{
			"4. three",
			[]int{1, 2, 3},
			1,
		},
		{
			"5. straight 3",
			[]int{3, 2, 1},
			3,
		},
		{
			"6. four numbers",
			[]int{3, 2, 1, 1},
			3,
		},
		{
			"7. five",
			[]int{3, 2, 1, 1, 2},
			3,
		},
		{
			"8. straight 3 2 1",
			[]int{3, 2, 1, 1, 1, 3, 2, 1},
			3,
		},
		{
			"9. long",
			[]int{12, 9, 10, 8, 6, 5, 3, 4, 2, 1},
			5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description,
			func(t *testing.T) {
				assert.Equal(t, tc.want, parse(tc.input))
			},
		)
	}
}
