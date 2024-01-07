package main

type letter string
type occurences int

func identifyLetters(word string) map[letter]occurences {
	result := make(map[letter]occurences)

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
