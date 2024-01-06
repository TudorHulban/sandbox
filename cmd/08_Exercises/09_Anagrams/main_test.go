package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnagramSlices(t *testing.T) {
	testCases := []struct {
		description string
		input1      []int
		input2      []int
		want        bool
	}{
		{
			"1. empty",
			[]int{},
			[]int{},
			true,
		},
		{
			"2. fail",
			[]int{1},
			[]int{},
			false,
		},
		{
			"3. fail",
			[]int{1},
			[]int{1, 2},
			false,
		},
		{
			"4. happy",
			[]int{-1},
			[]int{-1},
			true,
		},
		{
			"5. happy",
			[]int{-1, 2, -3},
			[]int{-3, -1, 2},
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description,
			func(t *testing.T) {
				assert.Equal(t,
					tc.want,
					checkAnagramSlices(
						tc.input1,
						tc.input2,
					),
				)
			},
		)
	}
}

func TestAnagramStrings(t *testing.T) {
	testCases := []struct {
		description string
		input1      string
		input2      string
		want        bool
	}{
		{
			"1. empty",
			"",
			"",
			true,
		},
		{
			"2. space",
			" ",
			" ",
			true,
		},
		{
			"3. japanese",
			"♥ab語 ",
			" b語a♥",
			true,
		},
		{
			"4. fail",
			"",
			"a",
			false,
		},
		{
			"5. fail",
			"a",
			"ab",
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description,
			func(t *testing.T) {
				assert.Equal(t,
					tc.want,
					checkAnagramStrings(
						tc.input1,
						tc.input2,
					),
				)
			},
		)
	}
}
