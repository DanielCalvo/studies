package main

import "fmt"

func Sum(i []int) (sum int) {
	for _, v := range i {
		sum += v
	}
	return sum
}

// Sums the contents of each slice passed and returns a slice with the results
// If it receives an empty slice, the sum is 0
func SumAll(numbersToSum ...[]int) (finalSlice []int) {
	for _, numberToSum := range numbersToSum {
		finalSlice = append(finalSlice, Sum(numberToSum))
	}
	return finalSlice
}

func SumAllTails(numbersToSum ...[]int) (finalSlice []int) {
	for _, numberToSum := range numbersToSum {
		if len(numberToSum) > 1 {
			finalSlice = append(finalSlice, Sum(numberToSum[1:]))
		}
	}
	return finalSlice
}

func main() {
	a := []int{1, 2}
	b := []int{3, 2}
	c := []int{4, 5, 6, 7}
	fmt.Println(SumAll(a, b, c))

	fmt.Println(SumAllTails([]int{}, []int{3}))

}
