package main

import (
	"fmt"
	"strings"
)

func main() {
	email := "john.doe@email.com"
	name, _ := getSenderFromEmail(email)

	fmt.Printf(
		"(%s)\n",
		strings.Join(name, ","),
	)
}
