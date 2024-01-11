package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMiddlePosition(t *testing.T) {
	testCases := []struct {
		description string
		input       string
		want        int
	}{
		{
			"1. aka",
			"aka",
			1,
		},
		{
			"2. abcd",
			"abcd",
			1,
		},
		{
			"3. abxcd",
			"abxcd",
			2,
		},
		{
			"4. abxycd",
			"abxycd",
			2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description,
			func(t *testing.T) {
				assert.Equal(t,
					tc.want,
					getMiddlePosition(tc.input),
				)
			},
		)
	}
}

func TestExtractLetter(t *testing.T) {
	testCases := []struct {
		description string
		input       string
		position    int
		want        string
	}{
		{
			"1. three letter word",
			"aka",
			1,
			"aa",
		},
		{
			"2. four letter word",
			"abcd",
			1,
			"acd",
		},
		{
			"3. six letter word",
			"abxycd",
			2,
			"abycd",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description,
			func(t *testing.T) {
				assert.Equal(t,
					tc.want,
					extractLetter(tc.input, tc.position),
				)
			},
		)
	}
}
