package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsStateValid(t *testing.T) {
	testCases := []struct {
		description string
		input       state
		want        bool
	}{
		{
			"1. nil",
			state{},
			true,
		},
		{
			"2. init",
			state{
				rodFrom: rod{3, 2, 1},
			},
			true,
		},
		{
			"3. happy 1",
			state{
				rodFrom: rod{3, 2},
				rodTo:   rod{1},
			},
			true,
		},
		{
			"4. happy 2",
			state{
				rodFrom: rod{3, 2},
				rodTo:   rod{1},
				rodAux:  rod{4},
			},
			true,
		},
		{
			"5. invalid",
			state{
				rodFrom: rod{3},
				rodTo:   rod{1, 4},
				rodAux:  rod{2},
			},
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description,
			func(t *testing.T) {
				assert.Equal(t,
					tc.want,
					tc.input.isValid(),
				)
			},
		)
	}
}
