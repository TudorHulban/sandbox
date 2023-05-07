package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConcatenates(t *testing.T) {
	x := []string{"1", "2", "3"}

	require.Equal(t, concatenateSlice(x), concatenateBuilder(x))
}

// cpu: AMD Ryzen 5 5600U with Radeon Graphics
// Benchmark_WSlice-12    	43538700	        32.33 ns/op	       3 B/op	       1 allocs/op
func Benchmark_WSlice(b *testing.B) {
	b.ResetTimer()

	x := []string{"1", "2", "3"}

	for i := 0; i < b.N; i++ {
		concatenateSlice(x)
	}
}

// cpu: AMD Ryzen 5 5600U with Radeon Graphics
// Benchmark_WBuilder-12    	48090902	        38.08 ns/op	       8 B/op	       1 allocs/op
func Benchmark_WBuilder(b *testing.B) {
	b.ResetTimer()

	x := []string{"1", "2", "3"}

	for i := 0; i < b.N; i++ {
		concatenateBuilder(x)
	}
}
