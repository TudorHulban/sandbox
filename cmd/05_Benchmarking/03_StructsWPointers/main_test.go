package main

import (
	"testing"
)

// cpu: AMD Ryzen 7 5700G with Radeon Graphics
// Benchmark_GetState_1-16         1000000000               0.2165 ns/op          0 B/op          0 allocs/op
// Benchmark_GetState_2-16         1000000000               0.6873 ns/op          0 B/op          0 allocs/op
// Benchmark_GetState_3-16         1000000000               0.2160 ns/op          0 B/op          0 allocs/op
// Benchmark_GetState_4-16         1000000000               0.2297 ns/op          0 B/op          0 allocs/op
// Benchmark_GetState_5-16         1000000000               0.4446 ns/op          0 B/op          0 allocs/op
// Benchmark_GetString-16          1000000000               0.2164 ns/op          0 B/op          0 allocs/op

func Benchmark_GetState_1(b *testing.B) {
	v := structWOPointers{
		state: "x",
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		getState1(v)
	}
}

func Benchmark_GetState_2(b *testing.B) {
	theState := "x"

	v := structWPointers{
		state: &theState,
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		getState2(v)
	}
}

func Benchmark_GetState_3(b *testing.B) {
	v := structWOPointers{
		state: "x",
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		getState3(&v)
	}
}

func Benchmark_GetState_4(b *testing.B) {
	v := structWOPointers{
		state: "x",
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		getState4(&v)
	}
}

func Benchmark_GetState_5(b *testing.B) {
	theState := "x"

	v := structWPointers{
		state: &theState,
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		getState5(&v)
	}
}

func Benchmark_GetString(b *testing.B) {
	v := "x"

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		getString(v)
	}
}
