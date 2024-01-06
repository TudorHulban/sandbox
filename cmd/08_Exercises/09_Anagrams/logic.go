package main

import "reflect"

func identifyLetters(text string) map[rune]int {
	res := make(map[rune]int)

	for _, letter := range text {
		if _, exists := res[letter]; exists {
			res[letter]++

			continue
		}

		res[letter] = 1
	}

	return res
}

func checkAnagramStrings(word1, word2 string) bool {
	return reflect.DeepEqual(
		identifyLetters(word1),
		identifyLetters(word2),
	)
}

func identifyElements[T comparable](slice []T) map[T]int {
	res := make(map[T]int)

	for _, element := range slice {
		if _, exists := res[element]; exists {
			res[element]++

			continue
		}

		res[element] = 1
	}

	return res
}

func checkAnagramSlices[T comparable](slice1, slice2 []T) bool {
	return reflect.DeepEqual(
		identifyElements(slice1),
		identifyElements(slice2),
	)
}
