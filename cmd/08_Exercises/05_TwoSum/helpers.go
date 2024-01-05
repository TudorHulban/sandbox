package main

type index int

func (ix index) asNumber() int {
	return int(ix)
}

func hashSlice(slice []int) map[int][]index {
	result := make(map[int][]index)

	for ix, value := range slice {
		if _, exists := result[value]; exists {
			result[value] = append(result[value],
				index(ix),
			)
		}

		result[value] = []index{
			index(ix),
		}
	}

	return result
}
