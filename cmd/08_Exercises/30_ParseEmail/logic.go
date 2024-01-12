package main

import "strings"

func getSenderFromEmail(email string) ([]string, error) {
	at := strings.Index(email, "@")
	if at == -1 {
		return nil,
			errMalformedEmail
	}

	name := strings.Split(email[:at], ".")
	if len(name) < 2 {
		return nil,
			errMalformedName
	}

	capSlice(name)

	return name, nil
}
