package main

import (
	"fmt"
	"math/rand/v2"
)

type testData struct {
	slice []int
	sum   int
}

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
	var myTestData []testData

	for i := 1; i < 100; i++ {
		fmt.Println("Loop:", 1)
		testSlice := GenerateSlice(i, i)
		fmt.Println(testSlice)

		testSum := GenerateRandomSum(testSlice)
		myTestData = append(myTestData, testData{slice: testSlice, sum: testSum})
		fmt.Println(testSlice)
	}

	fmt.Println(myTestData)

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

func GenerateRandomSum(s []int) int {
	if len(s) == 0 {
		return 0
	}

	if len(s) == 1 && s[0] == 0 {
		return 0
	}

	var pos1, pos2 int
	for pos1 == pos2 {
		pos1 = rand.IntN(len(s))
		pos2 = rand.IntN(len(s))
	}
	return s[pos1] + s[pos2]
}
