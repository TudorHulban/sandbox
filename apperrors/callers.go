package apperrors

import "strings"

type callers []string

func (c callers) String() string {
	result := make([]string, len(c))

	for i, j := len(c)-1, 0; i >= 0; i, j = i-1, j+1 {
		result[j] = c[i]
	}

	return strings.Join(result, ",")
}
