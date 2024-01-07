package main

import "strconv"

func computeSumOfDigits(n int) int {
	number := strconv.Itoa(n)

	if n < 0 {
		number = number[1:]
	}

	return sumDigits(number)
}

func conversionDigit(n string) int {
	res, _ := strconv.Atoi(n) // ignoring error handling for sake of reccursion

	return res
}

func sumDigits(number string) int {
	if len(number) == 1 {
		return conversionDigit(number)
	}

	return conversionDigit(number[:1]) + sumDigits(number[1:])
}
