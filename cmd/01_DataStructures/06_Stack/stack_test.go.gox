package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStack(t *testing.T) {
	var stackFolders stack

	stackFolders.push("folder1")
	stackFolders.push("subfolder")

	require.Equal(t, "folder1/subfolder", stackFolders.String())
}
