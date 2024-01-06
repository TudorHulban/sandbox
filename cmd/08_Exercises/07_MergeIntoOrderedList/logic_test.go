package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComparison(t *testing.T) {
	testCases := []struct {
		description string
		slice1      []int
		slice2      []int
		want        []int
	}{
		{
			"1. minimum arrays",
			[]int{1},
			[]int{2},
			[]int{1, 2},
		},
		{
			"2. minimum arrays",
			[]int{2},
			[]int{1},
			[]int{1, 2},
		},
		{
			"3. different sizes",
			[]int{1},
			[]int{3, 4},
			[]int{1, 3, 4},
		},
		{
			"4. different sizes",
			[]int{3, 4},
			[]int{1},
			[]int{1, 3, 4},
		},
		{
			"5. different positions",
			[]int{1, 3, 5, 7},
			[]int{2, 4, 6, 8, 9},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description,
			func(t *testing.T) {
				assert.Equal(t,
					tc.want,
					orderWComparison(tc.slice1, tc.slice2),
				)
			},
		)
	}
}

// cpu: AMD Ryzen 7 5800H with Radeon Graphics
// Benchmark_WPackSort-16           6392142               176.8 ns/op
// Benchmark_WComparison-16         5296520               286.8 ns/op

// os: red hat 9
// cpu: AMD Ryzen 7 5800H with Radeon Graphics
// Benchmark_WStandardLibSort-16           11751849                98.84 ns/op          104 B/op          2 allocs/op
// Benchmark_WComparison-16                 9190356               125.2 ns/op           248 B/op          5 allocs/op

func Benchmark_WStandardLibSort(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		orderWStandardLibSort(a1, a2)
	}
}

func Benchmark_WComparison(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		orderWComparison(a1, a2)
	}
}
