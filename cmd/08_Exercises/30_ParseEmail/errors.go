package main

import "errors"

var (
	errMalformedEmail = errors.New("email malformed")
	errMalformedName  = errors.New("name malformed")
)
