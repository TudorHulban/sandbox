package main

import (
	"fmt"
	"os"
)

func main() {
	word := "radkar"

	if isPalindrome(word) {
		fmt.Printf(
			"natural palindrome: %s",
			word,
		)

		os.Exit(0)
	}

	middlePos := getMiddlePosition(word)

	wordAdjusted := extractLetter(word, middlePos)

	if isPalindrome(wordAdjusted) {
		fmt.Printf(
			"helped palindrome: %s\n",
			wordAdjusted,
		)

		os.Exit(0)
	}

	fmt.Printf("not a palindrome: %s\n", word)
}
