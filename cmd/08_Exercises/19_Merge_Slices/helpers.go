package main

func isOrdered(slice []int) bool {
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			return false
		}
	}

	return true
}

func high(slice []int) int {
	return slice[len(slice)-1]
}

func low(slice []int) int {
	return slice[0]
}
