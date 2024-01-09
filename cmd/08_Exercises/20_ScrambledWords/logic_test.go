package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInScrambled(t *testing.T) {
	testCases := []struct {
		description string
		words       []string
		scrambled   string
		want        string
	}{
		{
			"1. errors - empty scrambled",
			[]string{"cat", "baby", "dog", "bird", "car"},
			"",
			"",
		},
		{
			"2. errors - one character",
			[]string{"cat"},
			"a",
			"",
		},
		{
			"3. errors - no found string",
			[]string{"a", "b", "c"},
			"d",
			"",
		},
		{
			"4. one character",
			[]string{"a"},
			"a",
			"a",
		},
		{
			"5. multiple words",
			[]string{"a", "b"},
			"b",
			"b",
		},
		{
			"6. various",
			[]string{"cat", "baby", "dog", "bird", "car"},
			"bbabylkkjabdog",
			"baby",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description,
			func(t *testing.T) {
				require.Equal(t,
					tc.want,
					parse(tc.scrambled, tc.words...),
				)
			},
		)
	}
}
