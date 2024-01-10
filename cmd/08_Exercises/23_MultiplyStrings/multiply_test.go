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
			"2. zero x zero",
			number("0"),
			number("0"),
			func(result int, errTest error) {
				require.NoError(t, errTest)
				require.Equal(t, 0, result)
			},
		},
		{
			"3. zero x one",
			number("0"),
			number("1"),
			func(resTest int, errTest error) {
				require.NoError(t, errTest)
				require.Equal(t, 0, resTest)
			},
		},
		{
			"4. valid multiplication",
			number("1"),
			number("2"),
			func(resTest int, errTest error) {
				require.NoError(t, errTest)
				require.Equal(t, 2, resTest)
			},
		},
		{
			"5. multi digit numbers",
			number("13"),
			number("20"),
			func(resTest int, errTest error) {
				require.NoError(t, errTest)
				require.Equal(t, 260, resTest)
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
