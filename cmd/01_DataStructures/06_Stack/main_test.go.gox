package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadFile(t *testing.T) {
	content, errRe := readFile(_filePath)
	require.NoError(t, errRe)

	fmt.Println(strings.Join(content, "\n"))
}

func TestFile(t *testing.T) {
	content, errRe := readFile(_filePath)
	require.NoError(t, errRe)

	entries := parse(convertToEntries(content))

	fmt.Println(strings.Join(entries, "\n"))
}
