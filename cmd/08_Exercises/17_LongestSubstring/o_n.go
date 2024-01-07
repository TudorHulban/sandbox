package main

type word []rune

func (t word) hasElement(el rune) bool {
	for _, letter := range t {
		if letter == el {
			return true
		}
	}

	return false
}

func (t *word) reset() {
	*t = []rune{}
}

func longestWord(word1, word2 word) word {
	if len(word1) > len(word2) {
		return word1
	}

	return word2
}

func parse(text string) string {
	var currentlongest, previousLongest word

	for _, element := range text {
		if currentlongest.hasElement(element) {
			if len(currentlongest) >= len(previousLongest) {
				previousLongest = append([]rune{}, currentlongest...)
			}

			currentlongest.reset()
		}

		currentlongest = append(currentlongest, element)
	}

	return string(longestWord(previousLongest, currentlongest))
}
