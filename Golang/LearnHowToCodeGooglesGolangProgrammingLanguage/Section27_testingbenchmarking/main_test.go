package main

import "testing"

func TestMultiplyByTen(t *testing.T) {
	if multiplyByTen(10) != 100 {
		t.Error("Expected 10, got:", multiplyByTen(10))
	}
}
