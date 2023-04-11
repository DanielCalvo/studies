package main

import (
	"testing"
)

type num struct {
	number int
	digits int
}

var testData = []num{
	{0, 1},
	{-0, 1},
	{-4, 1},
	{4, 1},
	{-42, 2},
	{42, 2},
	{-812, 3},
	{812, 3},
	{-5812, 4},
	{8512, 4},
	{-99812, 5},
	{15312, 5},
}

/*
DigitsInNumnber: 31.78 ns/op
CountDigits: 98.61 ns/op
Ha, take that ChatGPT!
*/

func BenchmarkDigitsInNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range testData {
			digits := DigitsInNumnber(v.number)
			if digits != v.digits {
				b.Fatalf("Got %d digits for %d", digits, v.digits)
			}
		}
	}
}

func BenchmarkCountDigits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range testData {
			digits := CountDigits(v.number)
			if digits != v.digits {
				b.Fatalf("Got %d digits for %d", digits, v.digits)
			}
		}
	}
}
