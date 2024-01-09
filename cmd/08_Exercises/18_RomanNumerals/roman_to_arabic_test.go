package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRomanToArabic(t *testing.T) {
	testCases := []struct {
		description string
		input       string
		check       func(int, error)
	}{
		{
			"1",
			"I",
			func(result int, err error) {
				require.NoError(t, err)
				require.Equal(t,
					1,
					result,
				)
			},
		},
	}

	for _, testCase := range testCases {
		testCase.check(
			romanToArabic(testCase.input),
		)
	}
}
