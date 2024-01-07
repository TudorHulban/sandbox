package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsRodValid(t *testing.T) {
	testCases := []struct {
		description string
		input       rod
		want        bool
	}{
		{
			"1. nil",
			rod{},
			true,
		},
		{
			"2. one disk",
			rod{1},
			true,
		},
		{
			"3. two disks",
			rod{2, 1},
			true,
		},
		{
			"4. three disks",
			rod{3, 2, 1},
			true,
		},
		{
			"5. invalid with two disks",
			rod{1, 2},
			false,
		},
		{
			"6. invalid with three disks",
			rod{2, 3, 1},
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
