package apperrors

import "fmt"

type ErrorValidation struct {
	Issue      error
	Resolution error

	CallerInitial     string
	CallersSubsequent callers
}

func (err ErrorValidation) Error() string {
	return fmt.Sprintf(
		"subsequentCallers: %s,\ninitialCaller: %s, issue: %s, resolution: %s",
		err.CallersSubsequent.String(),
		err.CallerInitial,
		err.Issue.Error(),
		err.Resolution,
	)
}

func (err *ErrorValidation) AddSubsequenCaller(callerName string) {
	err.CallersSubsequent = append(err.CallersSubsequent, callerName)
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
