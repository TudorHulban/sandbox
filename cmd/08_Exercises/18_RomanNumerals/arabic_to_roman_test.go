package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseArabicToRoman(t *testing.T) {
	testCases := []struct {
		description string
		input       uint
		want        string
	}{
		{
			"1",
			1,
			"I",
		},
		{
			"2",
			2,
			"II",
		},
		{
			"3",
			3,
			"III",
		},
		{
			"4",
			4,
			"IV",
		},
		{
			"8",
			8,
			"VIII",
		},
		{
			"9",
			9,
			"IX",
		},
		{
			"10",
			10,
			"X",
		},
		{
			"13",
			13,
			"XIII",
		},
		{
			"100",
			100,
			"C",
		},
	}

	for _, testCase := range testCases {
		roman := arabicToRoman(testCase.input)

		assert.Equal(t,
			testCase.want,
			roman,

			testCase.description,
		)
	}
}
