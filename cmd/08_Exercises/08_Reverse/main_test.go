package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseSliceInts(t *testing.T) {
	testCases := []struct {
		description string
		input       []int
		want        []int
	}{
		{
			"1. one element",
			[]int{7},
			[]int{7},
		},
		{
			"2. several elements",
			[]int{1, 2, 3, 4, 5},
			[]int{5, 4, 3, 2, 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description,
			func(t *testing.T) {
				assert.Equal(t,
					tc.want,
					reverseSlice(tc.input),
				)
			},
		)
	}
}

func TestReverseSliceStrings(t *testing.T) {
	testCases := []struct {
		description string
		input       []string
		want        []string
	}{
		{
			"1. one element",
			[]string{"a"},
			[]string{"a"},
		},
		{
			"2. several elements",
			[]string{"b", "b", "a", "a"},
			[]string{"a", "a", "b", "b"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description,
			func(t *testing.T) {
				assert.Equal(t,
					tc.want,
					reverseSlice(tc.input),
				)
			},
		)
	}
}

func TestReverseString(t *testing.T) {
	testCases := []struct {
		description string
		input       string
		want        string
	}{
		{
			"1. empty string",
			"",
			"",
		},
		{
			"2. one word",
			"abcde",
			"edcba",
		},
		{
			"3. japanese",
			"♥ab語",
			"語ba♥",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description,
			func(t *testing.T) {
				assert.Equal(t, tc.want, reverseString(tc.input))
			},
		)
	}
}
