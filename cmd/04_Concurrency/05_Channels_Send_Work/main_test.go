package main

import "testing"

func TestLogic(t *testing.T) {
	dispatcher(1, 10)
	dispatcher(3, 10)
	dispatcher(5, 10)
	dispatcher(8, 10)
	dispatcher(10, 10)
}
