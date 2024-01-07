package main

import "sort"

// MapData Exported for future use.
type MapString struct {
	Key   string
	Value int
}

func extractOccurences(letters map[letter]occurrences) []MapString {
	if len(letters) == 0 {
		return nil
	}

	result := make([]MapString, 0)

	for letter, occurrences := range letters {
		result = append(result,
			MapString{
				Key:   string(letter),
				Value: int(occurrences),
			},
		)
	}

	sort.Slice(result,
		func(i, j int) bool {
			return result[i].Value < result[j].Value
		},
	)

	return result
}
