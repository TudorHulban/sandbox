package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLongestWord(t *testing.T) {
	testCases := []struct {
		description string
		val         string
		expected    string
	}{
		{
			"1. repeating",
			"aa",
			"a",
		},
		{
			"2. simple",
			"aab",
			"ab",
		},
		{
			"3. complex",
			"abcabcbb",
			"abc",
		},
		{
			"4. substring",
			"aa12345bb",
			"a12345b",
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t,
			testCase.expected,
			parse(testCase.val),

			testCase.description,
		)
	}
}
