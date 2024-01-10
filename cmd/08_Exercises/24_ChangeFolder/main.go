package main

import (
	"strings"

	"fmt"
)

var currentFolder string

func main() {
	currentFolder = "/"

	input := make([]string, 0)

	input = append(input, "pwd")
	input = append(input, "cd /home/csed")
	input = append(input, "pwd")
	input = append(input, "cd x")
	input = append(input, "pwd")
	input = append(input, "cd ../..")
	input = append(input, "pwd") // should print /home

	process(input)
}

func process(inputs []string) {
	for _, input := range inputs {
		if input == "pwd" {
			fmt.Println("pwd:", currentFolder)

			continue
		}

		nextPath := strings.Split(input, " ")[1]

		currentFolder = handlePathLogic(nextPath)
	}
}
