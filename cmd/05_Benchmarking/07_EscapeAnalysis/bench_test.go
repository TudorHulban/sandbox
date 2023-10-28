package escapeanalysis

import (
	"fmt"
	"runtime"
	"testing"
)

var many = 10000

// with car as pointer
// cpu: AMD Ryzen 7 5700G with Radeon Graphics
// Benchmark_ages-16                       1000000000               0.0007543 ns/op               0 B/op          0 allocs/op
// Benchmark_agesPtr-16                    1000000000               0.0005545 ns/op               0 B/op          0 allocs/op
// Benchmark_agesPtrOptim-16               1000000000               0.0005552 ns/op               0 B/op          0 allocs/op
// Benchmark_agesResults-16                1000000000               0.0006596 ns/op               0 B/op          0 allocs/op
// Benchmark_agesResultsOptim-16           1000000000               0.0004468 ns/op               0 B/op          0 allocs/op

// with car as underlying type
// cpu: AMD Ryzen 7 5700G with Radeon Graphics
// Benchmark_ages-16                       1000000000               0.0009919 ns/op               0 B/op          0 allocs/op
// Benchmark_agesPtr-16                    1000000000               0.0003808 ns/op               0 B/op          0 allocs/op
// Benchmark_agesPtrOptim-16               1000000000               0.0003084 ns/op               0 B/op          0 allocs/op
// Benchmark_agesResults-16                1000000000               0.0003747 ns/op               0 B/op          0 allocs/op
// Benchmark_agesResultsOptim-16           1000000000               0.0002824 ns/op               0 B/op          0 allocs/op

func Benchmark_ages(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	var m1, m2 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m1)

	results := ages(uint(many))

	_ = results[len(results)-1].price

	for _, result := range results {
		_ = result.price
	}

	runtime.ReadMemStats(&m2)
	fmt.Println("execution total KB:", (m2.TotalAlloc-m1.TotalAlloc)/1024)
	fmt.Println("mallocs:", m2.Mallocs-m1.Mallocs)
}

func Benchmark_agesPtr(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	var m1, m2 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m1)

	results := agesPtr(uint(many))

	_ = results[len(results)-1].price

	for _, result := range results {
		_ = result.price
	}

	runtime.ReadMemStats(&m2)
	fmt.Println("execution total KB:", (m2.TotalAlloc-m1.TotalAlloc)/1024)
	fmt.Println("mallocs:", m2.Mallocs-m1.Mallocs)
}

func Benchmark_agesPtrOptim(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	var m1, m2 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m1)

	results := agesPtrOptim(uint(many))

	_ = results[len(results)-1].price

	for _, result := range results {
		_ = result.price
	}

	runtime.ReadMemStats(&m2)
	fmt.Println("execution total KB:", (m2.TotalAlloc-m1.TotalAlloc)/1024)
	fmt.Println("mallocs:", m2.Mallocs-m1.Mallocs)
}

func Benchmark_agesResults(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	var m1, m2 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m1)

	var results []*person

	agesResults(uint(many), &results)

	_ = results[len(results)-1].price

	for _, result := range results {
		_ = result.price
	}

	runtime.ReadMemStats(&m2)
	fmt.Println("execution total KB:", (m2.TotalAlloc-m1.TotalAlloc)/1024)
	fmt.Println("mallocs:", m2.Mallocs-m1.Mallocs)
}

func Benchmark_agesResultsOptim1(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	var m1, m2 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m1)

	results := make([]*person, many)

	agesResultsOptim(uint(many), &results)

	_ = results[len(results)-1].price

	for _, result := range results {
		_ = result.price
	}

	runtime.ReadMemStats(&m2)
	fmt.Println("execution total KB:", (m2.TotalAlloc-m1.TotalAlloc)/1024)
	fmt.Println("mallocs:", m2.Mallocs-m1.Mallocs)
}

func Benchmark_agesResultsOptim2(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	var m1, m2 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m1)

	results := make([]*person, many)

	agesResultsOptim(uint(many), &results)

	_ = results[len(results)-1].price

	for ix := range results {
		_ = results[ix].price
	}

	runtime.ReadMemStats(&m2)
	fmt.Println("execution total KB:", (m2.TotalAlloc-m1.TotalAlloc)/1024)
	fmt.Println("mallocs:", m2.Mallocs-m1.Mallocs)
}
