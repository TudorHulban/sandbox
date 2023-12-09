package main

import (
	"encoding/binary"
	"math/rand"
)

func generateSliceWRandom(size int) []int {
	b := []byte("always do good")

	random := rand.New( //nolint:gosec
		rand.NewSource(int64(binary.LittleEndian.Uint64(b))),
	)

	result := make([]int, size)

	for i := 0; i < size; i++ {
		result[i] = random.Intn(99) - rand.Intn(99) //nolint:gosec
	}

	return result
}
