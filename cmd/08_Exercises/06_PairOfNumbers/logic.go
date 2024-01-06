package main

func sumOfTwo(slice1, slice2 []int, sumAmount int) [][2]int {
	var result [][2]int

	for _, valueSlice1 := range slice1 {
		for _, valueSlice2 := range slice2 {
			if sumAmount == valueSlice1+valueSlice2 {
				result = append(result,
					[2]int{
						valueSlice1,
						valueSlice2,
					},
				)
			}
		}
	}

	return result
}
