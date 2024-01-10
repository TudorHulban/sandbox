package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConversion(t *testing.T) {
	testCases := []struct {
		description string
		input       string
		want        int
	}{
		{
			"0",
			"0",
			0,
		},
		{
			"1",
			"1",
			1,
		},
		{
			"11",
			"11",
			11,
		},
		{
			"13",
			"13",
			13,
		},
		{
			"20",
			"20",
			20,
		},
		{
			"111",
			"111",
			111,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description,
			func(t *testing.T) {
				require.Equal(t,
					tc.want,
					conversion(tc.input),
				)
			},
		)
	}
}
