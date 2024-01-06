package main

import (
	"testing"
)

const sum = 4

var a1 = []int{-1, -2, 4, 4, -2, -6, 5, -7}
var a2 = []int{0, 6, 3, 4, 0}

// RedHat 9
// cpu: AMD Ryzen 7 5800H with Radeon Graphics
// Benchmark_sumOfTwo-16    	 9179704	       132.7 ns/op	     240 B/op	       4 allocs/op
func Benchmark_sumOfTwo(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		sumOfTwo(a1, a2, sum)
	}
}
