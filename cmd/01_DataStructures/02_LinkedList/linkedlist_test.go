package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConstructor(t *testing.T) {
	list := newLinkedListFromHead(
		&node{
			data: 1,
		},
	)

	var buf bytes.Buffer

	fmt.Fprint(&buf, list)

	require.Equal(t, "1", buf.String())
}

func TestPrepend(t *testing.T) {
	list := newLinkedListFromHead(
		&node{
			data: 1,
		},
	)

	list.prepend(
		&node{
			data: 2,
		},

		&node{
			data: 3,
		},
	)

	var buf bytes.Buffer

	fmt.Fprint(&buf, list)

	require.Equal(t, "321", buf.String())
}

func TestDeleteNodeWValue(t *testing.T) {
	testCases := []struct {
		description string
		input       []*node
		value       int
		want        string
	}{
		{"head only value exists", []*node{{data: 1}}, 1, ""},
		{"head only value does not exist", []*node{{data: 1}}, 7, "1"},
		{"nodes value exists", []*node{{data: 1}, {data: 2}, {data: 3}}, 1, "32"},
		{"nodes value does not exist", []*node{{data: 1}, {data: 2}, {data: 3}}, 7, "321"},
	}

	for _, tc := range testCases {
		t.Run(tc.description,
			func(t *testing.T) {
				l := newLinkedListFromNodes(tc.input...)

				_ = l.deleteValue(tc.value)

				var buf bytes.Buffer
				fmt.Fprint(&buf, l)

				assert.Equal(t, tc.want, buf.String())
			},
		)
	}
}

func TestReverse(t *testing.T) {
	testCases := []struct {
		description string
		input       []*node
		want        string
	}{
		{"head only", []*node{{data: 1}}, "1"},
		{"nodes", []*node{{data: 1}, {data: 2}, {data: 3}}, "123"},
	}

	for _, tc := range testCases {
		t.Run(tc.description,
			func(t *testing.T) {
				l := newLinkedListFromNodes(tc.input...)
				l.reverse()

				var buf bytes.Buffer

				fmt.Fprint(&buf, l)

				assert.Equal(t, tc.want, buf.String())
			},
		)
	}
}
