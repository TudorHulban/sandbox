package main

func findValues(slice []int, sumAmount int) [2]int {
	hash := hashSlice(slice)

	for ix, value := range slice {
		difference := sumAmount - value

		if _, exists := hash[difference]; exists {
			for _, sliceIndex := range hash[difference] {
				if index(ix) != sliceIndex {
					return [2]int{
						ix,
						sliceIndex.asNumber(),
					}
				}
			}
		}
	}

	return [2]int{}
}
