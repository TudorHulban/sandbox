package main

import (
	"fmt"
	"regexp"

	"github.com/TudorHulban/GolangSandbox/helpers"
)

func returnID(url, regex string) (string, error) {
	r, errParse := regexp.Compile(regex)
	if errParse != nil {
		return "", errParse
	}

	id := helpers.High(r.FindStringSubmatch(url))

	return id, nil
}

func main() {
	id1, _ := returnID(
		"/id/123",
		"/id/(.*)",
	)

	fmt.Println("value:", id1)

	id2, _ := returnID(
		"/id/124/abc",
		"/id/124/(.*)",
	)

	fmt.Println("value:", id2)
}
