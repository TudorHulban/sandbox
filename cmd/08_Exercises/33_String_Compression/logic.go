package main

import "strconv"

func compress(word string) string {
	var res string

	var currentLetter string
	var currentNumOccurences int

	// no rune needed based on assumption.
	for i := 0; i < len(word); i++ {
		if i == 0 {
			currentLetter = word[0:1]
		}

		if currentLetter == word[i:i+1] {
			currentNumOccurences++
		}

		if currentLetter != word[i:i+1] {
			res = res + currentLetter + strconv.Itoa(currentNumOccurences)

			currentLetter = word[i : i+1]
			currentNumOccurences = 1
		}

		if i == len(word)-1 {
			res = res + currentLetter + strconv.Itoa(currentNumOccurences)
		}
	}

	if len(res) >= len(word) {
		return word
	}

	return res
}
