package main

import (
	"encoding/binary"
	"math/rand"
)

// generateSlice Generates a slice of passed size and filled with random numbers.
func generateSlice(size int) []int {
	b := []byte("always do good")

	random := rand.New(
		rand.NewSource(int64(binary.LittleEndian.Uint64(b))),
	)

	result := make([]int, size)

	for i := 0; i < size; i++ {
		result[i] = random.Intn(99) - rand.Intn(99)
	}

	return result
}
