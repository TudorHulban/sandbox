package main

// time complexity O(n)
// space complexity O(1)

// sort012 sorts a slice that contains numbers 0,1,2
func sort012(slice []int) []int {
	high := len(slice) - 1
	low := 0
	mid := 0

	for mid <= high {
		switch slice[mid] {
		case 0:
			slice[mid], slice[low] = slice[low], slice[mid]
			low++

		case 2:
			slice[mid], slice[high] = slice[high], slice[mid]
			high--

			mid--
		}

		mid++
	}

	return slice
}
