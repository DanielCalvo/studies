package main

import (
	"fmt"
	"math/rand/v2"
)

//This would be a cool one to benchmark!
//And also refresh your knowledge of execution times

func main() {
	slice := GenerateSlice(10, 10)
	sum := GenerateRandomSum(slice)
	fmt.Println("slice:", slice)
	fmt.Println("sum:", sum)
	pos := FindSumInSlice(slice, sum)
	fmt.Println("Found sum of slice at positions:", pos)

	//write a function for this so you can solve the leetcode!
	//benchmarking with your improved idea after you solved the leetcode would be cool too!

}

// I'd rather return (pos1, pos2 int) but problem wants []int

func FindSumInSliceFaster(slice []int, sum int) []int {

	return []int{}
}

func FindSumInSlice(slice []int, sum int) []int {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice); j++ {
			if slice[i]+slice[j] == sum {
				if i == j {
					continue
				}
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1} //No sum sound
}

func GenerateSlice(lenght, maxValue int) []int {
	s := make([]int, lenght)
	for i := 0; i < len(s); i++ {
		s[i] = rand.IntN(maxValue)
	}
	return s
}

func GenerateSliceTemplate

func GenerateRandomSum(s []int) int {
	var pos1, pos2 int
	for pos1 == pos2 {
		pos1 = rand.IntN(len(s))
		pos2 = rand.IntN(len(s))
	}
	return s[pos1] + s[pos2]
}
