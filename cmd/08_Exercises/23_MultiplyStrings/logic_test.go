package main

import (
	"testing"

	"github.com/TudorHulban/GolangSandbox/apperrors"
	"github.com/stretchr/testify/require"
)

func TestMultiply(t *testing.T) {
	testCases := []struct {
		description string
		inputA      number
		inputB      number
		checkResult func(int, error)
	}{
		{
			"1. errors - validation",
			number("o"),
			number("0"),
			func(result int, errTest error) {
				require.Error(t, errTest)
				require.ErrorAs(t,
					errTest,
					&apperrors.ErrorValidation{},
				)
				require.Zero(t, result)
			},
		},
		{
			"0",
			number("0"),
			number("0"),
			func(result int, errTest error) {
				require.NoError(t, errTest)
				require.Equal(t, 0, result)
			},
		},
		{
			"1",
			number("0"),
			number("1"),
			func(resTest int, errTest error) {
				require.NoError(t, errTest)
				require.Equal(t, 0, resTest)
			},
		},
		{
			"2",
			number("1"),
			number("2"),
			func(resTest int, errTest error) {
				require.NoError(t, errTest)
				require.Equal(t, 2, resTest)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description,
			func(t *testing.T) {
				tc.checkResult(
					multiply(tc.inputA, tc.inputB),
				)
			},
		)
	}
}
