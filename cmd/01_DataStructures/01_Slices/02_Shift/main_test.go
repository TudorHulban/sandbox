package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShiftRight(t *testing.T) {
	testCases := []struct {
		description string
		input       []int
		want        []int
	}{
		{"1", []int{1}, []int{1}},
		{"1,2", []int{1, 2}, []int{2, 1}},
		{"1,2,3", []int{1, 2, 3}, []int{3, 1, 2}},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			shiftRight(&tc.input)

			assert.Equal(t, tc.want, tc.input)
		})
	}
}
