package main

import (
	"math"
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
	if len(s) == 1 {
		return int([]rune(s)[0]) - int(rune('0'))
	}

	var sum int

	for power, pos := len(s), 0; pos < len(s); power, pos = power-1, pos+1 {
		sum = sum + int(math.Pow10(power-1))*int([]rune(s[pos : pos+1])[0]-rune('0'))
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
