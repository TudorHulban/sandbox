package escapeanalysis

import (
	"fmt"
	"runtime"
	"testing"
)

const withStats = false

var many = 256

// conclusion: results is better than pointer in most cases
// copy is better than results if the memory is available
// there is no best solution
// based on each situation an analysis should be done

// with car as pointer

// cpu: AMD Ryzen 7 5700G with Radeon Graphics
// Benchmark_ages-16                       1000000000               0.0007543 ns/op               0 B/op          0 allocs/op
// Benchmark_agesPtr-16                    1000000000               0.0005545 ns/op               0 B/op          0 allocs/op
// Benchmark_agesPtrOptim-16               1000000000               0.0005552 ns/op               0 B/op          0 allocs/op
// Benchmark_agesResults-16                1000000000               0.0006596 ns/op               0 B/op          0 allocs/op
// Benchmark_agesResultsOptim-16           1000000000               0.0004468 ns/op               0 B/op          0 allocs/op

// levelled:
// cpu: AMD Ryzen 7 5700G with Radeon Graphics
// Benchmark_ages-16                             74          29413740 ns/op        39500043 B/op     100029 allocs/op
// Benchmark_agesPtr-16                          61          21117873 ns/op        16901442 B/op     200028 allocs/op
// Benchmark_agesPtrOptim-16                     63          18683210 ns/op        13602858 B/op     200001 allocs/op
// Benchmark_agesResults-16                      76          20379528 ns/op        16901434 B/op     200028 allocs/op
// Benchmark_agesResultsOptim1-16                64          20326645 ns/op        13602883 B/op     200001 allocs/op
// Benchmark_agesResultsOptim2-16                74          18396600 ns/op        13602870 B/op     200001 allocs/op - best results for large sets

// cpu: AMD Ryzen 7 5700G with Radeon Graphics
// Benchmark_ages-16                        1000000              1300 ns/op            2720 B/op         15 allocs/op
// Benchmark_agesPtr-16                      721142              1774 ns/op            1528 B/op         25 allocs/op
// Benchmark_agesPtrOptim-16                 648300              1637 ns/op            1360 B/op         21 allocs/op
// Benchmark_agesResults-16                  626036              2028 ns/op            1528 B/op         25 allocs/op
// Benchmark_agesResultsOptim1-16           1009758              1303 ns/op            1360 B/op         21 allocs/op - best results for small sets
// Benchmark_agesResultsOptim2-16            651938              1565 ns/op            1360 B/op         21 allocs/op

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

	if withStats {
		runtime.GC()
		runtime.ReadMemStats(&m1)
	}

	for i := 0; i < b.N; i++ {
		results := ages(uint(many))

		_ = results[len(results)-1].price

		for _, result := range results {
			_ = result.price
		}
	}

	if withStats {
		runtime.ReadMemStats(&m2)
		fmt.Println("execution total KB:", (m2.TotalAlloc-m1.TotalAlloc)/1024)
		fmt.Println("mallocs:", m2.Mallocs-m1.Mallocs)
	}
}

func Benchmark_agesPtr(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	var m1, m2 runtime.MemStats

	if withStats {
		runtime.GC()
		runtime.ReadMemStats(&m1)
	}

	for i := 0; i < b.N; i++ {
		results := agesPtr(uint(many))

		_ = results[len(results)-1].price

		for _, result := range results {
			_ = result.price
		}
	}

	if withStats {
		runtime.ReadMemStats(&m2)
		fmt.Println("execution total KB:", (m2.TotalAlloc-m1.TotalAlloc)/1024)
		fmt.Println("mallocs:", m2.Mallocs-m1.Mallocs)
	}
}

func Benchmark_agesPtrOptim(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	var m1, m2 runtime.MemStats

	if withStats {
		runtime.GC()
		runtime.ReadMemStats(&m1)
	}

	for i := 0; i < b.N; i++ {
		results := agesPtrOptim(uint(many))

		_ = results[len(results)-1].price

		for _, result := range results {
			_ = result.price
		}
	}

	if withStats {
		runtime.ReadMemStats(&m2)
		fmt.Println("execution total KB:", (m2.TotalAlloc-m1.TotalAlloc)/1024)
		fmt.Println("mallocs:", m2.Mallocs-m1.Mallocs)
	}
}

func Benchmark_agesResults(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	var m1, m2 runtime.MemStats

	if withStats {
		runtime.GC()
		runtime.ReadMemStats(&m1)
	}

	for i := 0; i < b.N; i++ {
		var results []*person

		agesResults(uint(many), &results)

		_ = results[len(results)-1].price

		for _, result := range results {
			_ = result.price
		}
	}

	if withStats {
		runtime.ReadMemStats(&m2)
		fmt.Println("execution total KB:", (m2.TotalAlloc-m1.TotalAlloc)/1024)
		fmt.Println("mallocs:", m2.Mallocs-m1.Mallocs)
	}
}

func Benchmark_agesResultsOptim1(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	var m1, m2 runtime.MemStats

	if withStats {
		runtime.GC()
		runtime.ReadMemStats(&m1)
	}

	for i := 0; i < b.N; i++ {
		results := make([]*person, many, many)

		agesResultsOptim(uint(many), &results)

		_ = results[len(results)-1].price

		for _, result := range results {
			_ = result.price
		}
	}

	if withStats {
		runtime.ReadMemStats(&m2)
		fmt.Println("execution total KB:", (m2.TotalAlloc-m1.TotalAlloc)/1024)
		fmt.Println("mallocs:", m2.Mallocs-m1.Mallocs)
	}
}

func Benchmark_agesResultsOptim2(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	var m1, m2 runtime.MemStats

	if withStats {
		runtime.GC()
		runtime.ReadMemStats(&m1)
	}

	for i := 0; i < b.N; i++ {
		results := make([]*person, many, many)

		agesResultsOptim(uint(many), &results)

		_ = results[len(results)-1].price

		for ix := range results {
			_ = results[ix].price
		}
	}

	if withStats {
		runtime.ReadMemStats(&m2)
		fmt.Println("execution total KB:", (m2.TotalAlloc-m1.TotalAlloc)/1024)
		fmt.Println("mallocs:", m2.Mallocs-m1.Mallocs)
	}
}
