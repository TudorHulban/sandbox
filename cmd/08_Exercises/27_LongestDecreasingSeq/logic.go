package main

import (
	"github.com/TudorHulban/GolangSandbox/helpers"
)

func parse(slice []int) int {
	var previousNumber int

	var previousRange int

	var currentRange int

	for ix, number := range slice {
		if ix == 0 {
			previousNumber = number
			previousRange++
			currentRange++

			continue
		}

		if previousNumber > number {
			previousNumber = number
			currentRange++

			if ix == len(slice)-1 {
				return helpers.Max(currentRange, previousRange)
			}

			continue
		}

		previousRange = helpers.Max(currentRange, previousRange)

		previousNumber = number

		currentRange = 1
	}

	return previousRange
}
