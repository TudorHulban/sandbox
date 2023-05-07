package main

import "testing"

// cpu: AMD Ryzen 5 5600U with Radeon Graphics
// BenchmarkWType-12    	1000000000	         0.0000002 ns/op	       0 B/op	       0 allocs/op
func BenchmarkWType(b *testing.B) {
	a := []int{1, 2, 3, 4, 5}

	b.ResetTimer()
	delIndex(a, 3)
}

// cpu: AMD Ryzen 5 5600U with Radeon Graphics
// BenchmarkWGenerics-12    	1000000000	         0.0000001 ns/op	       0 B/op	       0 allocs/op
func BenchmarkWGenerics(b *testing.B) {
	a := []int{1, 2, 3, 4, 5}

	b.ResetTimer()
	delIndexG(a, 3)
}
