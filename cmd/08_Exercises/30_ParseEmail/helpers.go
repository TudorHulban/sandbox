package main

import "strings"

func capFirstLetter(text string) string {
	if len(text) == 0 {
		return ""
	}

	word := []rune(text)

	return strings.ToUpper(string(word[0])) + string(word[1:])
}

func capSlice(slice []string) {
	for i := 0; i < len(slice); i++ {
		slice[i] = capFirstLetter(slice[i])
	}
}
