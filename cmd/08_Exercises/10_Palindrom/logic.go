package main

func isPalindrome(text string) bool {
	word := []rune(text)

	if len(word) == 1 {
		return true
	}

	if len(word) == 2 {
		return text[0] == text[1]
	}

	if len(word) == 3 {
		return text[0] == text[2]
	}

	half := len(word) / 2

	h1 := word[:half]
	h2 := word[half+1:]

	for i := 0; i < half-1; i++ {
		if h1[i] != h2[len(h2)-1-i] {
			return false
		}
	}

	return true
}
