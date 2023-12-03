package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddNumberValue(t *testing.T) {
	persons := []person{
		*newPerson("john", 1, 2),
		*newPerson("mary", 1, 2),
	}
	updateByValue(persons, 3)

	require.Equal(t, []int{1, 2}, persons[0].numbers)
	require.Equal(t, []int{1, 2}, persons[1].numbers)
}

func TestAddNumberReference(t *testing.T) {
	persons := []person{
		*newPerson("john", 1, 2),
		*newPerson("mary", 1, 2),
	}
	updateByReference(&persons, 3)

	require.Equal(t, []int{1, 2, 3}, persons[0].numbers)
	require.Equal(t, []int{1, 2, 3}, persons[1].numbers)
}

// cpu: AMD Ryzen 7 5800H with Radeon Graphics
// BenchmarkAddNumberReference-16    	76035600	        14.88 ns/op	      96 B/op	       0 allocs/op
func BenchmarkAddNumberReference(b *testing.B) {
	persons := []person{
		*newPerson("john", 1, 2),
		*newPerson("mary", 1, 2),
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		updateByReference(&persons, 3)
	}
}

func TestAddNumberPointers(t *testing.T) {
	persons := []*person{
		newPerson("john", 1, 2),
		newPerson("mary", 1, 2),
	}
	updateByPointers(persons, 3)

	require.Equal(t, []int{1, 2, 3}, persons[0].numbers)
	require.Equal(t, []int{1, 2, 3}, persons[1].numbers)
}

// cpu: AMD Ryzen 7 5800H with Radeon Graphics
// BenchmarkAddNumberPointers-16    	94505680	        16.00 ns/op	      97 B/op	       0 allocs/op
func BenchmarkAddNumberPointers(b *testing.B) {
	persons := []*person{
		newPerson("john", 1, 2),
		newPerson("mary", 1, 2),
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		updateByPointers(persons, 3)
	}
}

func TestAddNumberPointersSugar(t *testing.T) {
	persons := []*person{
		newPerson("john", 1, 2),
		newPerson("mary", 1, 2),
	}
	updatePreferred(persons, 3)

	require.Equal(t, []int{1, 2, 3}, persons[0].numbers)
	require.Equal(t, []int{1, 2, 3}, persons[1].numbers)
}

// cpu: AMD Ryzen 7 5800H with Radeon Graphics
// BenchmarkAddNumberPreferred-16    	75272113	        14.23 ns/op	      97 B/op	       0 allocs/op
func BenchmarkAddNumberPreferred(b *testing.B) {
	persons := []*person{
		newPerson("john", 1, 2),
		newPerson("mary", 1, 2),
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		updatePreferred(persons, 3)
	}
}
