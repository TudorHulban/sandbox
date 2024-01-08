package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	cases := []struct {
		description string

		s1, s2 []int
		want   []int
	}{
		{"case 1", []int{0, 2, 4}, []int{1, 3}, []int{0, 1, 2, 3, 4}},
		{"case 2", []int{0, 1, 2}, []int{1, 3, 4}, []int{0, 1, 1, 2, 3, 4}},
	}

	for _, tc := range cases {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.want, merge(tc.s1, tc.s2))
		})
	}
}
