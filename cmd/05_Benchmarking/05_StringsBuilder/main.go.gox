package main

import (
	"strings"
)

func main() {

}

func concatenateSlice(s []string) string {
	return strings.Join(s, "")
}

func concatenateBuilder(s []string) string {
	var b strings.Builder

	for _, value := range s {
		b.WriteString(value)
	}

	return b.String()
}
