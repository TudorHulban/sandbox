package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringCompression(t *testing.T) {
	testCases := []struct {
		description string
		input       string
		want        string
	}{
		{
			"1",
			"aabcccccaaa",
			"a2b1c5a3",
		},
		{
			"2",
			"aa",
			"aa",
		},
		{
			"3",
			"aabbb",
			"a2b3",
		},
		{
			"4",
			"aaAAA",
			"a2A3",
		},
		{
			"5",
			"aaA",
			"aaA",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description,
			func(t *testing.T) {
				assert.Equal(t,
					tc.want,
					compress(tc.input),
				)
			},
		)
	}
}
