package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var exEmptySlice = []string{}
var exNonEmptySlice = []string{"1", "2", "3", "4", "5"}

func Test_Nil(t *testing.T) {
	assert.True(t, byLength(exEmptySlice))
}

// Benchmark_Len-4   	1000000000	         0.767 ns/op	       0 B/op	       0 allocs/op
func Benchmark_Len1(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		byLength(exEmptySlice)
	}
}

// Benchmark_Nil-4   	1000000000	         0.764 ns/op	       0 B/op	       0 allocs/op
func Benchmark_Nil1(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		byNil(exEmptySlice)
	}
}

// Benchmark_Len-4   	1000000000	         0.767 ns/op	       0 B/op	       0 allocs/op
func Benchmark_Len2(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		byLength(exNonEmptySlice)
	}
}

// Benchmark_Nil-4   	1000000000	         0.764 ns/op	       0 B/op	       0 allocs/op
func Benchmark_Nil2(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		byNil(exNonEmptySlice)
	}
}

// cpu: AMD Ryzen 7 5700G with Radeon Graphics
// Benchmark_Len1-16       1000000000               0.2168 ns/op          0 B/op          0 allocs/op
// Benchmark_Nil1-16       1000000000               0.2303 ns/op          0 B/op          0 allocs/op
// Benchmark_Len2-16       1000000000               0.2159 ns/op          0 B/op          0 allocs/op
// Benchmark_Nil2-16       1000000000               0.2300 ns/op          0 B/op          0 allocs/op
