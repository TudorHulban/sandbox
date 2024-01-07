package main

type letter string
type occurrences int

func identifyLetters(word string) map[letter]occurrences {
	result := make(map[letter]occurrences)

	for i := 0; i < len(word); i++ {
		element := word[i : i+1]

		if _, exists := result[letter(element)]; exists {
			result[letter(element)]++

			continue
		}

		result[letter(element)] = 1
	}

	return result
}
