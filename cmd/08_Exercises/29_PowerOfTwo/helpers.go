package main

import "strconv"

func integerToBinary(number int) string {
	return strconv.FormatInt(int64(number), 2)
}
