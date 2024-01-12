package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCapFirstLetter(t *testing.T) {
	testCases := []struct {
		description string
		input       string
		want        string
	}{
		{
			"1. errors - one number",
			"1",
			"1",
		},
		{
			"2. errors - empty",
			"",
			"",
		},
		{
			"3. one leter",
			"a",
			"A",
		},
		{
			"4. two leters",
			"at",
			"At",
		},
		{
			"5. UTF no cap variation",
			"♥ymy♥",
			"♥ymy♥",
		},
		{
			"6. UTF with cap variation",
			"āt",
			"Āt",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description,
			func(t *testing.T) {
				assert.Equal(t,
					tc.want,
					capFirstLetter(tc.input),
				)
			},
		)
	}
}
