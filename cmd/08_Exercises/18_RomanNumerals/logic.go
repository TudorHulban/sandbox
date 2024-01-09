package main

import (
	"fmt"
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
				mapArabicToRoman[int(conversionValues[i])],
			)
		}

		buffer = buffer % conversionValues[i]
	}

	return strings.Join(result, "")
}

func romanToArabic(number string) (int, error) {
	var result int

	for pos := 0; pos < len(number); pos++ {
		if pos < len(number)-1 {
			arabic2Pos, exists2Pos := mapRomanToArabic[number[pos:pos+2]]
			if exists2Pos {
				result = result + int(arabic2Pos)

				if pos == len(number)-2 {
					break
				}

				pos++

				continue
			}
		}

		arabic1Pos, exists1Pos := mapRomanToArabic[number[pos:pos+1]]
		if exists1Pos {
			result = result + int(arabic1Pos)

			continue
		}

		return 0,
			fmt.Errorf(
				"not found in conversion table: %s",
				number[pos:pos+1],
			)
	}

	return result,
		nil
}
