package main

import "sort"

// MapData Exported for future use.
type MapString struct {
	Key   string
	Value int
}

func extractOccurences(letters map[letter]occurences) []MapString {
	if len(letters) == 0 {
		return nil
	}

	var result []MapString

	for letter, occurences := range letters {
		result = append(result,
			MapString{
				Key:   string(letter),
				Value: int(occurences),
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
