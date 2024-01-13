package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEvaluate(t *testing.T) {
	testCases := []struct {
		description string
		input       string
		check       func(int, error)
	}{
		{
			"1. errors - unknown operation",
			"2*6",
			func(result int, err error) {
				require.Error(t, err)
				require.Zero(t, result)
			},
		},
		{
			"2. 2+6+4",
			"2+6+4",
			func(result int, err error) {
				require.NoError(t, err)
				require.Equal(t,
					12,
					result,
				)
			},
		},
		{
			"3. 20-10-6",
			"20-10-6",
			func(result int, err error) {
				require.NoError(t, err)
				require.Equal(t,
					4,
					result,
				)
			},
		},
		{
			"4. 123+200",
			"123+200",
			func(result int, err error) {
				require.NoError(t, err)
				require.Equal(t,
					323,
					result,
				)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description,
			func(t *testing.T) {
				tc.check(
					evaluate(tc.input),
				)
			},
		)
	}
}
