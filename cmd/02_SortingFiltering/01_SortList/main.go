package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

const (
	sortLength = iota
	sortAlphabetic
)

var sortMethod int64

func main() {
	flagSort := flag.String("sort", "0", "sort method, 0 - by length, 1 - alphabetic")
	flag.Parse()

	sortMethod, _ = strconv.ParseInt(*flagSort, 10, 64)
	if (sortMethod > 1) || (sortMethod < 0) {
		fmt.Printf("Bad parameter: %d", sortMethod)

		os.Exit(1)
	}

	method := []string{"by Length", "Alphabetic"}

	log.Println("Flag sort: ", sortMethod, "(", method[sortMethod], ")")

	data := []string{"xen", "zi", "ab", "a"}

	sort.Sort(JustWords(data))

	log.Println(data)
}

//https://golang.org/pkg/sort/
