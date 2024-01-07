package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLinkedList_Constructor(t *testing.T) {
	l1 := NewLinkedList(
		&Node{
			Value: 1,
		},
	)

	number, errGet := l1.GetNumber()
	require.NoError(t, errGet)
	require.Equal(t,
		1,
		number,
	)
}

func TestLinkedList_Nil(t *testing.T) {
	l1 := NewLinkedList(
		&Node{
			Value: 1,
		},
	)

	l1.Prepend(
		&Node{
			Value: 2,
		},
	)

	number, errGet := l1.GetNumber()
	require.NoError(t, errGet)
	require.Equal(t,
		21,
		number,
	)
}

func TestNewLinkedListFromNumber(t *testing.T) {
	ll := NewLinkedListFromNumber(1)
	require.NotNil(t, ll)

	listNumber, errNumber := ll.GetNumber()
	require.NoError(t, errNumber)
	require.Equal(t,
		1,
		listNumber,
	)
}
