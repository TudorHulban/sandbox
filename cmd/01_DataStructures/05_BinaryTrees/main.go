package main

import (
	"fmt"
)

func main() {
	n := newNode(10)
	n.Add(12)
	n.Add(3)
	n.Add(11)

	fmt.Println(n)

	key := 12
	var retrieved node

	n.GetNode(key, &retrieved)

	fmt.Println(retrieved.Info())

	fmt.Printf("has key %d - %t\n", key, n.HasKey(key))
}
