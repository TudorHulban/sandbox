package main

func isPalindrome(text string) bool {
	word := []rune(text)

	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
		if word[i] != word[j] {
			return false
		}
	}

	return true
}
