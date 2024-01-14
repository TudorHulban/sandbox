package main

import "strconv"

func compress(word string) string {
	var result string

	var currentLetter string

	var currentNumOccurences int

	// no rune needed based on assumption.
	for i := 0; i < len(word); i++ {
		if i == 0 {
			currentLetter = word[0:1]
		}

		if currentLetter == word[i:i+1] {
			currentNumOccurences++
		} else {
			result = result + currentLetter + strconv.Itoa(currentNumOccurences)

			currentLetter = word[i : i+1]
			currentNumOccurences = 1
		}

		if i == len(word)-1 {
			result = result + currentLetter + strconv.Itoa(currentNumOccurences)
		}
	}

	if len(result) >= len(word) {
		return word
	}

	return result
}
