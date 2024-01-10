package helpers

type Index int

func (ix Index) AsNumber() int {
	return int(ix)
}

func HashSlice[T comparable](slice []T) map[T][]Index {
	result := make(map[T][]Index)

	for ix, value := range slice {
		if _, exists := result[value]; exists {
			result[value] = append(result[value],
				Index(ix),
			)
		}

		result[value] = []Index{
			Index(ix),
		}
	}

	return result
}
