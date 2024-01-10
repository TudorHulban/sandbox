package apperrors

import "fmt"

type ErrorValidation struct {
	Issue error

	Caller     string
	Resolution string
}

func (err ErrorValidation) Error() string {
	return fmt.Sprintf(
		"caller: %s, issue: %s, resolution: %s",
		err.Caller,
		err.Issue.Error(),
		err.Resolution,
	)
}

func (err *ErrorValidation) AddResolution(reso error) {
	err.Resolution = reso.Error()
}

func (err *ErrorValidation) OverwriteCaller(newCaller string) {
	err.Caller = newCaller
}

type ErrorInvalidInput struct {
	InputName string
}

func (err ErrorInvalidInput) Error() string {
	return fmt.Sprintf(
		"invalid input name: '%s'",
		err.InputName,
	)
}
