package main

// os: mint 21.3
// cpu: AMD Ryzen 7 5800H with Radeon Graphics
// BenchmarkGORMSQLite-16              9616            131452 ns/op            5556 B/op        130 allocs/op
// BenchmarkGORMPG-16                  3483            453224 ns/op            5587 B/op         98 allocs/op
// BenchmarkGORMRAWPG-16               3799            340385 ns/op            3807 B/op         77 allocs/op
// BenchmarkPGXSimple-16               5822            216842 ns/op             560 B/op         12 allocs/op
// BenchmarkPGXPool-16                 5097            239600 ns/op             722 B/op         13 allocs/op
