package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	got := Sum(numbers)
	want := 15
	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

func TestSumAll(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("Check if function works", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{3, 6})
		want := []int{3, 9}
		checkSums(t, got, want)
	})
	t.Run("Check if function works with 1 element", func(t *testing.T) {
		got := SumAll([]int{2}, []int{6})
		want := []int{2, 6}
		checkSums(t, got, want)
	})

	t.Run("Trying to sum empty slices", func(t *testing.T) {
		got := SumAll([]int{}, []int{})
		want := []int{0, 0}
		checkSums(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2, 3}, []int{0, 9, 1})
	want := []int{5, 10}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

/*
func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Joe", "")
		want := "Hello, Joe"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})
}
*/
