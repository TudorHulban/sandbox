package main

import (
	"strings"
)

// Implements the sort interface.
type JustWords []string

func (w JustWords) Len() int {
	return len(w)
}

func (w JustWords) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func (w JustWords) Less(i, j int) bool {
	switch sortMethod {
	case sortLength:
		return len(w[i]) < len(w[j])

	case sortAlphabetic:
		return strings.ToLower(w[i]) < strings.ToLower(w[j])
	}

	return false
}
