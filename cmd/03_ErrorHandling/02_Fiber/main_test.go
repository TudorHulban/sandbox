package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// see https://pkg.go.dev/errors#Is
func TestErrorHandlingBadInput(t *testing.T) {
	errInput := validateInput(badInput)

	require.ErrorIs(t, errInput, errorBadInput, errInput)
}

// see https://pkg.go.dev/errors#As
func TestErrorHandlingBadText(t *testing.T) {
	errInput := validateInput(badText)

	var errText *errorBadText

	require.ErrorAs(t, errInput, &errText, errInput)
}

func TestErrorHandlingBadSpelling(t *testing.T) {
	errInput := validateInput(badSpelling)

	require.ErrorIs(t, errInput, errorBadText{badText: badSpelling}, errInput)
	require.ErrorAs(t, errInput, &errorBadText{badText: badSpelling}, errInput)
}
