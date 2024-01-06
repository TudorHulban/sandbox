package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPalindrome(t *testing.T) {
	testCases := []struct {
		description string
		word        string
		want        bool
	}{
		{
			"1. one letter",
			"a",
			true,
		},
		{
			"2. fail on two letters",
			"ab",
			false,
		},
		{
			"3. pass on two letters",
			"aa",
			true,
		},
		{
			"4. fail on three letters",
			"abc",
			false,
		},
		{
			"5. pass on three letters",
			"aba",
			true,
		},
		{
			"6. pass on five",
			"tenet",
			true,
		},
		{
			"7. pass on unicode",
			"♥ymy♥",
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description,
			func(t *testing.T) {
				assert.True(t,
					isPalindrome(tc.word) == tc.want,

					"result of palindrome: '%t' for: '%s'",
					isPalindrome(tc.word),
					tc.word,
				)
			},
		)
	}
}
