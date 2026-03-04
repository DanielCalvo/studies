package main

//on your test files, remember to import the test package!
import "testing"

// lets test the sum function -- idiomatic go is to call it "Test" followed by your function name, in camelcase!
func TestMysum(t *testing.T) {
	got := mysum(2, 3)
	want := 5
	//the got and want variables are very common in go, they're used in the standard library too
	if got != want {
		t.Fatalf("mysum test failed, got: %d, want: %d", got, want)
	}
}
