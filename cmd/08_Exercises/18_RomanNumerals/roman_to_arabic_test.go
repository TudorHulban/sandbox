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
			"error character",
			"i",
			func(result int, err error) {
				require.Error(t, err)
				require.Empty(t,
					result,
				)
			},
		},
		{
			"error typo",
			"Vi",
			func(result int, err error) {
				require.Error(t, err)
				require.Empty(t,
					result,
				)
			},
		},
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
		{
			"2",
			"II",
			func(result int, err error) {
				require.NoError(t, err)
				require.Equal(t,
					2,
					result,
				)
			},
		},
		{
			"3",
			"III",
			func(result int, err error) {
				require.NoError(t, err)
				require.Equal(t,
					3,
					result,
				)
			},
		},
		{
			"4",
			"IV",
			func(result int, err error) {
				require.NoError(t, err)
				require.Equal(t,
					4,
					result,
				)
			},
		},
		{
			"8",
			"VIII",
			func(result int, err error) {
				require.NoError(t, err)
				require.Equal(t,
					8,
					result,
				)
			},
		},
		{
			"9",
			"IX",
			func(result int, err error) {
				require.NoError(t, err)
				require.Equal(t,
					9,
					result,
				)
			},
		},
		{
			"401",
			"CDI",
			func(result int, err error) {
				require.NoError(t, err)
				require.Equal(t,
					401,
					result,
				)
			},
		},
		{
			"998",
			"CMXCVIII",
			func(result int, err error) {
				require.NoError(t, err)
				require.Equal(t,
					998,
					result,
				)
			},
		},
		{
			"1111",
			"MCXI",
			func(result int, err error) {
				require.NoError(t, err)
				require.Equal(t,
					1111,
					result,
				)
			},
		},
		{
			"2222",
			"MMCCXXII",
			func(result int, err error) {
				require.NoError(t, err)
				require.Equal(t,
					2222,
					result,
				)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description,
			func(t *testing.T) {
				testCase.check(
					romanToArabic(testCase.input),
				)
			},
		)
	}
}
