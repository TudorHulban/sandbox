package main

// time complexity O(n)
// space complexity O(1)

// sort01 sorts a slice that contains numbers 0,1
func sort01(slice []int) []int {
	high := len(slice) - 1
	low := 0

	for low <= high {
		switch slice[low] {
		case 1:
			slice[low], slice[high] = slice[high], slice[low]
			high--
		}

		low++
	}

	return slice
}
