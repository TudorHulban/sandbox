package main

func mapStrings(slice []string) map[string]int {
	result := make(map[string]int)

	for _, word := range slice {
		if _, exists := result[word]; exists {
			result[word]++

			continue
		}

		result[word] = 1
	}

	return result
}

func mapsIncluded(shell, includes map[string]int) bool {
	if len(shell) < len(includes) {
		return false
	}

	for word := range includes {
		if _, exists := shell[word]; !exists {
			return false
		}

		if shell[word] < includes[word] {
			return false
		}
	}

	return true
}

func sliceIncluded(slice, includes []string) bool {
	return mapsIncluded(
		mapStrings(slice),
		mapStrings(includes),
	)
}
