package main

import "github.com/TudorHulban/GolangSandbox/helpers"

func findValues(slice []int, sumAmount int) [2]int {
	hash := helpers.HashSlice(slice)

	for ix, value := range slice {
		difference := sumAmount - value

		if _, exists := hash[difference]; exists {
			for _, sliceIndex := range hash[difference] {
				if helpers.Index(ix) != sliceIndex {
					return [2]int{
						ix,
						sliceIndex.AsNumber(),
					}
				}
			}
		}
	}

	return [2]int{}
}
