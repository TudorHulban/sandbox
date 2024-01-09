package main

func contains(word, shell string) bool {
	for ix := range shell {
		if ix+len(word) > len(shell) {
			return false
		}

		if word == shell[ix:ix+len(word)] {
			return true
		}
	}

	return false
}

func parse(scrambled string, lookFor ...string) string {
	for _, word := range lookFor {
		if contains(word, scrambled) {
			return word
		}
	}

	return ""
}
