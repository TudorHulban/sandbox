package main

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseBruteInts(t *testing.T) {
	testCases := []struct {
		description string
		a1          []int
		a2          []int
		want        []int
	}{
		{
			"1. no overlap a1 on left",
			[]int{1, 2},
			[]int{5, 6, 13, 14},
			[]int{1, 2, 5, 6, 13, 14},
		},
		{
			"2. overlap middle",
			[]int{5, 6},
			[]int{1, 2, 13, 14},
			[]int{1, 2, 5, 6, 13, 14},
		},
		{
			"3. overlap left",
			[]int{1, 6},
			[]int{2, 5, 13, 14},
			[]int{1, 2, 5, 6, 13, 14},
		},
		{
			"4. overlap right",
			[]int{5, 14},
			[]int{1, 2, 6, 13},
			[]int{1, 2, 5, 6, 13, 14},
		},
		{
			"5. no overlap a1 on right",
			[]int{13, 14},
			[]int{1, 2, 5, 6},
			[]int{1, 2, 5, 6, 13, 14},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description,
			func(t *testing.T) {
				ordered := parseBrute(
					tt.a1,
					tt.a2,
				)

				assert.Equal(t,
					tt.want,
					ordered,

					tt.description,
				)
				assert.True(t,
					slices.IsSorted(ordered),

					"slice not ordered",
				)
			},
		)
	}
}

func TestParseBruteUInts(t *testing.T) {
	testCases := []struct {
		description string
		a1          []uint
		a2          []uint
		want        []uint
	}{
		{
			"1. uint",
			[]uint{1, 2, 3},
			[]uint{5, 6, 7, 8},
			[]uint{1, 2, 3, 5, 6, 7, 8},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.description,
			func(t *testing.T) {
				ordered := parseBrute(
					tt.a1,
					tt.a2,
				)

				assert.Equal(t,
					tt.want,
					ordered,

					tt.description,
				)
				assert.True(t,
					slices.IsSorted(ordered),

					"slice not ordered",
				)
			},
		)
	}
}
