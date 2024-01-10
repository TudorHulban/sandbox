package main

import (
	"strconv"

	"github.com/TudorHulban/GolangSandbox/apperrors"
)

type number string

func (n number) IsNumber() error {
	_, errConv := strconv.Atoi(string(n))

	if errConv != nil {
		return apperrors.ErrorValidation{
			Caller: "IsNumber",
			Issue:  errConv,
		}
	}

	return nil
}

func conversion(s string) int {
	// TODO: add input validation by comparing to rune of zero

	var sum int

	for i := len(s) - 1; i >= 0; i-- {
		var buf int

		if i == 0 {
			buf = int(rune(s[i]) - rune('0'))
		}

		if i > 0 {
			buf = i * 10 * int(rune(s[i])-rune('0'))
		}

		sum = sum + buf
	}

	return sum
}

func multiply(a, b number) (int, error) {
	if errValidation := a.IsNumber(); errValidation != nil {
		errResult := errValidation.(apperrors.ErrorValidation)
		errResult.OverwriteCaller("multiply")
		errResult.AddResolution(
			apperrors.ErrorInvalidInput{
				InputName: "a",
			},
		)

		return 0,
			errResult
	}

	if errValidation := b.IsNumber(); errValidation != nil {
		errResult := errValidation.(apperrors.ErrorValidation)
		errResult.OverwriteCaller("multiply")
		errResult.AddResolution(
			apperrors.ErrorInvalidInput{
				InputName: "a",
			},
		)

		return 0,
			errResult
	}

	return conversion(string(a)) * conversion(string(b)), nil
}
