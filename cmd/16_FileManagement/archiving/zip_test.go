package archiving

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestZip(t *testing.T) {
	var z Zip

	fileName := "./constants.go"

	require.NoError(t,
		z.CompressFile(fileName),
	)
	require.True(t,
		fileExists(fileName+zipExtension),
	)
}
