package main

import (
	"fmt"
	"os"
)

func main() {
	var files []string

	if errWalk := doPathWalk(thePath, &files); errWalk != nil {
		fmt.Println(errWalk)

		os.Exit(1)
	}

	for _, file := range files {
		fmt.Println(file)
	}
}
