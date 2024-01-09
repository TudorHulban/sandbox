package main

import (
	"strings"
)

func arabicToRoman(number uint) string {
	var result []string

	buffer := number

	for i := len(conversionValues) - 1; i >= 0; i-- {
		remainder := buffer / conversionValues[i]

		if remainder == 0 {
			continue
		}

		for ix := 0; ix < int(remainder); ix++ {
			result = append(result,
				conversionMap[int(conversionValues[i])],
			)
		}

		buffer = buffer % conversionValues[i]
	}

	return strings.Join(result, "")
}

func romanToArabic(number string) (int, error) {
	var result int

	for pos := range number {

	}

	return result,
		nil
}
