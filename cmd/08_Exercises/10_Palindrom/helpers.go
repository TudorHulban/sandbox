package main

import "strings"

func getMiddlePosition(word string) int {
	var res int

	for i, j := 0, 0; j < len(word); i, j = i+1, j+2 {
		res = i
	}

	return res
}

func extractLetter(word string, position int) string {
	var res []string

	for i := 0; i < len(word); i = i + 1 {
		if i == position {
			continue
		}

		res = append(res, word[i:i+1])
	}

	return strings.Join(res, "")
}
