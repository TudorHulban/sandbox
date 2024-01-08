package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBrute(t *testing.T) {
	a1 := []int{1, 2, 3}
	a2 := []int{5, 6, 7, 8}
	a3 := []int{4}

	m := mergeBrute(a1, a2, a3)
	t.Log(m)

	require.True(t, isOrdered(m))
}
