package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPowerOfTwo(t *testing.T) {
	testCases := []struct {
		description string
		input       int
		want        bool
	}{
		{
			"negative number",
			-1,
			false,
		},
		{
			"0",
			0,
			false,
		},
		{
			"1",
			1,
			true,
		},
		{
			"2",
			2,
			true,
		},
		{
			"3",
			3,
			false,
		},
		{
			"5",
			5,
			false,
		},
		{
			"64",
			64,
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description,
			func(t *testing.T) {
				assert.Equal(t,
					tc.want,
					isPowerOfTwo(tc.input),
				)
			},
		)
	}
}
