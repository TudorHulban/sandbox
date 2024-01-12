package main

// time complexity: O(1)
// space complexity: O(1)

func isPowerOfTwo(number int) bool {
	if number <= 0 {
		return false
	}

	return number&(number-1) == 0
}
