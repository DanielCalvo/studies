package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestGenerateTestData(t *testing.T) {
	var myTestData []testData

	for i := 1; i < 100; i++ {
		testSlice := GenerateSlice(i, i)
		testSum := GenerateRandomSum(testSlice)
		myTestData = append(myTestData, testData{slice: testSlice, sum: testSum})
		fmt.Println(testSlice)
	}

	fmt.Println(myTestData)

}

func TestGenerateSlice(t *testing.T) {
	tests := []struct {
		length   int
		maxValue int
	}{
		{length: 10, maxValue: 100},
		{length: 1, maxValue: 1},
		{length: 2, maxValue: 2},
		{length: 0, maxValue: 100},
		{length: 100, maxValue: 1},
		{length: 1000, maxValue: 1000000},
	}

	for _, test := range tests {
		slice := GenerateSlice(test.length, test.maxValue)
		if len(slice) != test.length {
			t.Errorf("Generated slice length is incorrect. Expected %d, got %d", test.length, len(slice))
		}
		for _, value := range slice {
			if value < 0 || value >= test.maxValue {
				t.Errorf("Generated slice contains out-of-range value: %d", value)
			}
		}
	}
}

func TestFindSumInSlice(t *testing.T) {
	tests := []struct {
		slice      []int
		sum        int
		wantResult []int
	}{
		{slice: []int{1, 2, 3, 6}, sum: 5, wantResult: []int{1, 2}},
		{slice: []int{0, 0, 0, 0}, sum: 0, wantResult: []int{0, 1}},
		{slice: []int{4, 4, 4}, sum: 8, wantResult: []int{0, 1}},
		{slice: []int{1, 2, 3, 4}, sum: 20, wantResult: []int{-1, -1}}, //if sum is not found, return [-1,-1]
	}
	for _, test := range tests {
		result := FindSumInSlice(test.slice, test.sum)
		if !slices.Equal(test.wantResult, result) {
			t.Errorf("Result is incorrect, expected %d, got %d, input slice: %d", test.wantResult, result, test.slice)
		}
	}
}

func TestGenerateRandomSum(t *testing.T) {
	tests := []struct {
		slice []int
		sum   int
	}{
		{slice: []int{}, sum: 0},
		{slice: []int{0}, sum: 0},
		{slice: []int{1, 1}, sum: 2},
		{slice: []int{1, 2}, sum: 3},
	}
	for _, test := range tests {
		result := GenerateRandomSum(test.slice)
		if result != test.sum {
			t.Errorf("Result is incorrect, expected %d, got %d", test.sum, result)
		}
	}
}

func BenchmarkFindSumInSlice(b *testing.B) {
	// Generate a random slice and sum for benchmarking
	slice := GenerateSlice(10, 10)
	sum := GenerateRandomSum(slice)

	// Run the benchmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindSumInSlice(slice, sum)
	}
}
