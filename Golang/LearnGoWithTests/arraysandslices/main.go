package main

import "fmt"

func Sum(i []int) (sum int) {
	for _, v := range i {
		sum += v
	}
	return sum
}

// Receives: []int{1,2} and []int{4,5}, returns []int{3,9}
func SumAll(numbersToSum ...[]int) (finalSlice []int) {
	for _, numberToSum := range numbersToSum {
		finalSlice = append(finalSlice, Sum(numberToSum))
	}
	return finalSlice
}

func main() {
	a := []int{1, 2}
	b := []int{3, 2}
	fmt.Println(SumAll(a, b))
}
