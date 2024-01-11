package helpers

func Closer(from1, number, from2 uint) (uint, uint) {
	diffFrom1 := number - from1
	diffFrom2 := from2 - number

	if diffFrom1 > diffFrom2 {
		return from2, diffFrom2
	}

	return from1, diffFrom1
}

func Max[T uint | int](number1, number2 T) T {
	if number1 > number2 {
		return number1
	}

	return number2
}
