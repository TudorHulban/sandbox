package main

func mapGeneric[T string | int](slice []T) map[T]int {
	result := make(map[T]int)

	for _, word := range slice {
		if _, exists := result[word]; exists {
			result[word]++

			continue
		}

		result[word] = 1
	}

	return result
}

func mapsIncluded[T string | int](shell, includes map[T]int) bool {
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

func sliceIncluded[T string | int](slice, includes []T) bool {
	return mapsIncluded(
		mapGeneric(slice),
		mapGeneric(includes),
	)
}
