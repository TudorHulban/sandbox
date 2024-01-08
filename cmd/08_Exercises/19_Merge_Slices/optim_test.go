package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge2(t *testing.T) {
	testCases := []struct {
		description string
		a1          []int
		a2          []int
		want        []int
	}{
		{"no overlap a1 on left", []int{1, 2}, []int{5, 6, 13, 14}, []int{1, 2, 5, 6, 13, 14}},
		{"overlap middle", []int{5, 6}, []int{1, 2, 13, 14}, []int{1, 2, 5, 6, 13, 14}},
		{"overlap left", []int{1, 6}, []int{2, 5, 13, 14}, []int{1, 2, 5, 6, 13, 14}},
		{"overlap right", []int{5, 14}, []int{1, 2, 6, 13}, []int{1, 2, 5, 6, 13, 14}},
		{"no overlap a1 on right", []int{13, 14}, []int{1, 2, 5, 6}, []int{1, 2, 5, 6, 13, 14}},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			m := merge2(tt.a1, tt.a2)

			assert.Equal(t, tt.want, m, tt.description)
			assert.True(t, isOrdered(m), "slice not ordered")
		})
	}
}
