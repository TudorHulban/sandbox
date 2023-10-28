package main

// go test -bench=.

import (
	"testing"
)

var sample []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

func benchmarkSumLoop(ints []int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		sumLoop(ints)
	}
}

func BenchmarkSumLoop(b *testing.B) {
	benchmarkSumLoop(sample, b)
}

func benchmarkSumRecursive(ints []int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		sumRecurs(ints)
	}
}

func BenchmarkSumRecurs(b *testing.B) {
	benchmarkSumRecursive(sample, b)
}

// cpu: AMD Ryzen 5 5600U with Radeon Graphics
// BenchmarkSumLoop-12      	229665572	         5.365 ns/op	       0 B/op	       0 allocs/op
// BenchmarkSumRecurs-12    	38141940	        31.55 ns/op	       0 B/op	       0 allocs/op

// cpu: AMD Ryzen 3 4300GE with Radeon Graphics
// BenchmarkSumLoop-8   	216120318	         5.515 ns/op	       0 B/op	       0 allocs/op
// BenchmarkSumRecurs-8   	44978752	        26.67 ns/op	       0 B/op	       0 allocs/op

// cpu: AMD Ryzen 7 5700G with Radeon Graphics
// BenchmarkSumRecurs-16    	52434220	        22.86 ns/op	       0 B/op	       0 allocs/op
