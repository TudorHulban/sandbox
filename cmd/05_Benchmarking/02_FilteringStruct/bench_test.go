package main

// go test -bench=.

import (
	"testing"
)

var fieldName = "Name"

func BenchmarkFilterDirect10K(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		work := NewTasks(10000)

		showFieldDirect(fieldName, work)
	}
}

func BenchmarkFilterDirect100K(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		work := NewTasks(100000)

		showFieldDirect(fieldName, work)
	}
}

func BenchmarkFilterAssertion10K(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		work := NewTasks(10000)

		showFieldAssertion(fieldName, work)
	}
}

func BenchmarkFilterAssertion100K(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		work := NewTasks(100000)

		showFieldAssertion(fieldName, work)
	}
}

func BenchmarkFilterName10K(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		work := NewTasks(10000)

		showFieldName("Name", work)
	}
}

func BenchmarkFilterName100K(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		work := NewTasks(100000)

		showFieldName("Name", work)
	}
}

// --------------- Test Results

// goos: linux
// goarch: amd64
// pkg: github.com/TudorHulban/GolangSandbox/28_Benchmarking/02_FilteringStruct
// cpu: AMD Ryzen 5 3400G with Radeon Vega Graphics
// BenchmarkDirect-8            277           4281532 ns/op
// BenchmarkAssert-8            283           4227647 ns/op
// BenchmarkName-8              426           2744627 ns/op
// PASS
// ok      github.com/TudorHulban/GolangSandbox/28_Benchmarking/02_FilteringStruct 4.704s
