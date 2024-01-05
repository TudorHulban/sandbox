package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSumOrdered(t *testing.T) {
	testCases := []struct {
		description string
		input       []int
		sumAmount   int
		want        [2]int
	}{
		{"1, 3, 4, 5, 7, 11", []int{1, 3, 4, 5, 7, 11}, 9, [2]int{4, 5}},
	}

	for _, tc := range testCases {
		t.Run(tc.description,
			func(t *testing.T) {
				assert.Equal(t,
					tc.want,
					twoSumOrdered(tc.input, tc.sumAmount),
				)
			},
		)
	}
}
