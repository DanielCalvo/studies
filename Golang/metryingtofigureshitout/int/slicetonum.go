//Given a number, say 12555 create a function that will return a slice of all the digits, ex: []int{1,2,5,5,5}
//And also, create a function that takes a slice of int (ex: []int{1,2,5,5,5}) and return it as a single integer (12555)

package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(slicetoNum([]int{1, 2, 3, 4}))

	fmt.Println(numToSlice(-12345))

	//s[1] = n % 10
	//n = n / 10
	//s[0] = n % 10

}

// take first element, add to int
// multiply int by 10
// add second element
// multiply int by 10
// add third element
// and so on until you're out of elements!
func slicetoNum(i []int) int {
	result := i[0]
	for j := 1; j < len(i); j++ {
		result = result * 10
		result = result + i[j]
	}
	return result
}

// Hang on -- what if the number is negative?
func numToSlice(n int) []int {
	var s []int
	for n != 0 {
		s = slices.Insert(s, 0, n%10)
		n = n / 10
	}
	return s
}
