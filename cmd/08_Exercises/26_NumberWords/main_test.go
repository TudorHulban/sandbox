package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestNoWords(t *testing.T) {
	testCases := []struct {
		description string
		input       string
		want        int
	}{
		{
			"1. empty",
			"",
			0,
		},
		{
			"2. separator",
			"  ",
			0,
		},
		{
			"3. one word",
			"a",
			1,
		},
		{
			"4. one word",
			"abc",
			1,
		},
		{
			"5. one word with separator on right",
			"a ",
			1,
		},
		{
			"6. one word with two separators",
			" a ",
			1,
		},
		{
			"7. one word with separator on left",
			" a",
			1,
		},
		{
			"8. two words",
			"a b",
			2,
		},
		{
			"9. two words",
			"a b ",
			2,
		},
		{
			"10. two words with separator",
			"a   b",
			2,
		},
		{
			"11. three words",
			"a bc d",
			3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description,
			func(t *testing.T) {
				assert.Equal(t,
					tc.want,
					noWords(tc.input),

					tc.description,
				)
			},
		)
	}
}
