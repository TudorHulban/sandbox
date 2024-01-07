package main

type letter string

// deleteOne Deletes recursively a letter from given string
// and adds combinations to result pointer.
func deleteOne(text string, res *map[letter]bool) {
	if len(text) == 1 {
		(*res)[letter(text)] = true

		return
	}

	for i := 0; i < len(text); i++ {
		element := text[:i] + text[i+1:]

		(*res)[letter(element)] = true

		deleteOne(element, res)
	}
}
