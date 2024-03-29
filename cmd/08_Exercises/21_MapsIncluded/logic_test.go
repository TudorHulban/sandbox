package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSliceIncludesStrings(t *testing.T) {
	testCases := []struct {
		description string
		shell       []string
		includes    []string
		want        bool
	}{
		{
			"1. errors - empty shell",
			nil,
			[]string{"a"},
			false,
		},
		{
			"2. shell shorter than includes",
			[]string{"a"},
			[]string{"a", "b"},
			false,
		},
		{
			"3. shell does not contain",
			[]string{"a", "b"},
			[]string{"c"},
			false,
		},
		{
			"4. shell contains includes",
			[]string{"a", "a"},
			[]string{"a"},
			true,
		},
		{
			"5. shell equals includes",
			[]string{"a"},
			[]string{"a"},
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description,
			func(t *testing.T) {
				require.Equal(t,
					tc.want,
					sliceIncluded(tc.shell, tc.includes),
				)
			},
		)
	}
}

func TestSliceIncludesInt(t *testing.T) {
	testCases := []struct {
		description string
		shell       []int
		includes    []int
		want        bool
	}{
		{
			"1. errors - empty shell",
			nil,
			[]int{1},
			false,
		},
		{
			"2. shell contains includes",
			[]int{1, 2},
			[]int{1},
			true,
		},
		{
			"5. shell equals includes",
			[]int{1},
			[]int{1},
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description,
			func(t *testing.T) {
				require.Equal(t,
					tc.want,
					sliceIncluded(tc.shell, tc.includes),
				)
			},
		)
	}
}
