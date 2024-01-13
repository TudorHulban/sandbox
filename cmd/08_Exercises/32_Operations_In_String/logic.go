package main

import "strconv"

type operation func(a, b int) int

var addition operation = func(a, b int) int {
	return a + b
}

var subtraction operation = func(a, b int) int {
	return a - b
}

func evaluate(exp string) (int, error) {
	operations := splitASCIIString(exp, []string{"+", "-"})

	oper := addition

	var result int

	for _, item := range operations {
		switch item {
		case "+":
			oper = addition

		case "-":
			oper = subtraction

		default:
			number, errConvert := strconv.Atoi(item)
			if errConvert != nil {
				return 0, errConvert
			}

			result = oper(result, number)
		}
	}

	return result, nil
}
