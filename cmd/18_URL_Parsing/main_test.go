package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewURL(t *testing.T) {
	url, errURL := NewURL("id", "123")
	require.NoError(t, errURL)
	require.NotZero(t, url)

	fmt.Println(url)
}
