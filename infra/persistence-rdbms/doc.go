package main

// os: mint 21.3
// cpu: AMD Ryzen 7 5800H with Radeon Graphics
// BenchmarkGORMSQLite-16         131452 ns/op            5556 B/op        130 allocs/op
// BenchmarkGORMPG-16             453224 ns/op            5587 B/op         98 allocs/op
// BenchmarkGORMRAWPG-16          340385 ns/op            3807 B/op         77 allocs/op
// BenchmarkPGXSimple-16          216842 ns/op             560 B/op         12 allocs/op
// BenchmarkPGXPool-16            239600 ns/op             722 B/op         13 allocs/op

// pgx insert example as per
// https://stackoverflow.com/questions/2944297/postgresql-function-for-last-inserted-id

// pgx function query example as per
// https://medium.com/geekculture/work-with-go-postgresql-using-pgx-caee4573672
