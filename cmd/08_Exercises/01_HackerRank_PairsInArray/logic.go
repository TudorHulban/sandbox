package main

import (
	"fmt"
	"log"
)

func getPairs(occurencesForPair int, slice []int) (int, error) {
	if occurencesForPair < 1 {
		return 0,
			fmt.Errorf(
				"passed occurences for pair number value: %v smaller than one occurence",
				occurencesForPair)
	}

	log.Printf("computing number of pairs for %v in pair.", occurencesForPair)

	// defining new map to hold the number of occurences per each value
	// key is value, value is number of occurences
	data := make(map[int]int)

	for _, item := range slice {
		if _, exists := data[item]; exists {
			data[item]++

			continue
		}

		data[item] = 1
	}

	log.Println("data:", data)

	// extract number of pairs
	// defining holder for the pairs
	// key is the number for which occurences are calculated and value the number of occurences
	pairs := make(map[int]int)

	for item, numberOccurences := range data {
		numberPairs := numberOccurences / occurencesForPair

		log.Printf("item: %d with occurences: %d and number of pairs: %d\n",
			item, numberOccurences, numberPairs)

		if numberPairs >= 1 {
			log.Printf("item: %d adding for number of pairs: %d",
				item, numberPairs)

			pairs[item] = numberPairs
		}
	}

	log.Println("pairs:", pairs)

	var result int

	for item, numberPairs := range pairs {
		fmt.Printf("adding for item: %d pairs: %d\n", item, numberPairs)

		result = result + numberPairs
	}

	return result, nil
}
