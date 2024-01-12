package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSenderFromEmail(t *testing.T) {
	testCases := []struct {
		description string
		input       string
		want        []string
		err         error
	}{
		{
			"1. errors - email malformed",
			"john.doe#email.com",
			nil,
			errMalformedEmail,
		},
		{
			"2. errors - name malformed 1",
			"johndoe@email.com",
			nil,
			errMalformedName,
		},
		{
			"3. errors - name malformed 2",
			"john,doe@email.com",
			nil,
			errMalformedName,
		},
		{
			"4. happy path 1",
			"john.doe@email.com",
			[]string{"John", "Doe"},
			nil,
		},
		{
			"5. happy path 2",
			"john.doe.junior@email.com",
			[]string{"John", "Doe", "Junior"},
			nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description,
			func(t *testing.T) {
				sender, errGetSender := getSenderFromEmail(tc.input)

				if tc.err != nil {
					assert.Equal(t, tc.err, errGetSender)

					return
				}

				assert.Equal(t, tc.want, sender)
			},
		)
	}
}
