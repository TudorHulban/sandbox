package escapeanalysis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ages(t *testing.T) {
	persons := ages(1)

	require.NotZero(t, persons[0])
}
