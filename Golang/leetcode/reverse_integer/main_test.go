package main

import "testing"

/*
Uh, what happens with -0? or -1?
If its a single digit number, return the number

*/

type num struct {
	got  int
	want int
}

var testData = []num{
	{0, 0},
	{-0, -0},
	{-4, -4},
	{4, 4},
	{-42, -24},
	{42, 24},
	{-812, -218},
	{812, 218},
	{800, 8},
	{1534236469, 9646324351},
}

func BenchmarkReverseNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range testData {
			digits := ReverseNumber(v.got)
			if digits != v.want {
				b.Fatalf("Got %d digits for %d", digits, v.want)
			}
		}
	}
}

func BenchmarkReverseNumberString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range testData {
			digits, err := ReverseNumberString(v.got)
			if err != nil {
				b.Error("Error on ReverseNumberString:", err)
			}
			if digits != v.want {
				b.Fatalf("Got %d digits for %d", digits, v.want)
			}
		}
	}
}
