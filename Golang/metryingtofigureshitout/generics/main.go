package main

import "fmt"

/*
 */

func main() {
	fmt.Println("Hello world!")

	ints := make(map[string]int64)
	ints["a"] = 1
	ints["b"] = 2
	fmt.Println(ints)
	fmt.Println(SumInts(ints))

	floats := map[string]float64{
		"z": 22.22,
		"x": 33.33,
	}
	fmt.Println(SumFloats(floats))
	fmt.Println(SumIntsOrFloats(ints))

	tt := []int{11, 12}
	fmt.Println(Sum(tt))

}

type Number interface {
	int | int64 | float64
}

func Sum[T Number](numbers []T) T {
	var total T
	for _, x := range numbers {
		total += x
	}
	return total
}

// This will sum either a bunch of ints or a bunch of float64s, but not both at the same time, hmm
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}
