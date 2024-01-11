package main

func main() {}

func isSeparator(letter rune) bool {
	return string(letter) == " "
}

func noWords(text string) int {
	var wordStarted bool

	var result int

	for _, letter := range text {
		if isSeparator(letter) {
			if wordStarted {
				result++

				wordStarted = false

				continue
			}

			continue
		}

		wordStarted = true
	}

	if wordStarted {
		return result + 1
	}

	return result
}
