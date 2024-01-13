package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestASCIISplit(t *testing.T) {
	testCases := []struct {
		description string
		input       string
		splitters   []string
		want        []string
	}{
		{
			"1+",
			"1+",
			[]string{"+", "-"},
			[]string{"1", "+"},
		},
		{
			"2+6",
			"2+6",
			[]string{"+", "-"},
			[]string{"2", "+", "6"},
		},
		{
			"2+6+4",
			"2+6+4",
			[]string{"+", "-"},
			[]string{"2", "+", "6", "+", "4"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description,
			func(t *testing.T) {
				assert.Equal(t,
					tc.want,
					splitASCIIString(tc.input, tc.splitters),
				)
			},
		)
	}
}
