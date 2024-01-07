package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStackBasicOps(t *testing.T) {
	var s stack

	el := 1

	s.push(el)
	require.Equal(t,
		el,
		s.peek(),
	)
	require.False(t,
		s.isEmpty(),
	)

	popedEl := s.pop()
	require.True(t,
		s.isEmpty(),
	)
	require.Equal(t,
		el,
		popedEl,
	)
}

func TestReverse(t *testing.T) {
	var s stack

	s.push(1)
	s.push(2)

	require.Equal(t,
		"21",
		s.String(),
	)

	s.reverse()

	require.Equal(t,
		"12",
		s.String(),
	)
}
