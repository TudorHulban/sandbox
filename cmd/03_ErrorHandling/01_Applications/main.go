package main

import (
	"errors"
	"fmt"
	"time"
)

func fooService() error {
	return &ErrService{
		Caller:     "fooService",
		MethodName: "m1",
		Issue:      errors.New("issue 1"),

		WithTime:            true,
		NanosecondsDuration: time.Since(time.Now()).Nanoseconds(),
	}
}

func fooDomain() error {
	return &ErrDomain{
		Caller:     "fooDomain",
		MethodName: "fooService",
		Issue:      fooService(),
	}
}

func main() {
	fmt.Println(fooDomain())
	// go run .

	// Area: Domain
	// Caller: fooDomain
	// Method Name: fooService
	// Issue:
	// Area: Service
	// Caller: fooService
	// Method Name: m1
	// Issue: issue 1
}
