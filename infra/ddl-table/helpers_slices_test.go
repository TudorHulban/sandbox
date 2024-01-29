package ddltable

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPop(t *testing.T) {
	s1 := []string{"1"}

	s1 = Pop[string](s1)
	require.Empty(t, s1)

	s2 := []string{"1", "2"}

	s2 = Pop[string](s2)
	require.Len(t, s2, 1)
}
