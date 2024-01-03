package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInSequence(t *testing.T) {
	testCases := []struct {
		description string
		input       []int
		sequence    []int
		wantError   bool
	}{
		{
			"1. fail fast 1",
			[]int{1},
			[]int{1, 2, 3},
			true,
		},
		{
			"2. fail incomplete",
			[]int{1, 2, 3, 1, 2, 3, 1},
			[]int{1, 2, 3},
			true,
		},
		{
			"3. fail no sequence",
			[]int{1, 2, 3, 1, 2, 4},
			[]int{1, 2, 3},
			true,
		},
		{
			"4. happy path",
			[]int{1, 2, 3, 1, 2, 3},
			[]int{1, 2, 3},
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description,
			func(t *testing.T) {
				err := inSequence(tc.input, tc.sequence)

				if !tc.wantError {
					assert.NoError(t, err)

					return
				}

				assert.Error(t, err)
			},
		)
	}
}
