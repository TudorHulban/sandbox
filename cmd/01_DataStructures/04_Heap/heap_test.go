package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHeap(t *testing.T) {
	g1 := data{
		priority: 1,
	}

	g2 := data{
		priority: 7,
	}

	g3 := data{
		priority: 3,
	}

	h := NewQueue()
	h.
		Insert(&g1).
		Insert(&g2).
		Insert(&g3)

	require.Len(t, h.state, 3)
}

func TestExtractNil(t *testing.T) {
	h := NewQueue()

	require.Nil(t, h.Extract())
}

func TestExtractAll(t *testing.T) {
	g1 := data{
		priority: 1,
	}

	h := NewQueue()
	h.Insert(&g1).Extract()

	require.Len(t, h.state, 0)
}

func TestExtract(t *testing.T) {
	g1 := data{
		priority: 1,
	}

	g2 := data{
		priority: 7,
	}

	g3 := data{
		priority: 3,
	}

	g4 := data{
		priority: 5,
	}

	h := NewQueue()
	h.
		Insert(&g1).
		Insert(&g2).
		Insert(&g3).
		Insert(&g4)

	require.Equal(t, h.Extract().priority, g2.priority)
	require.Len(t, h.state, 3)

	require.Equal(t, h.Extract().priority, g4.priority)
	require.Len(t, h.state, 2)
}
