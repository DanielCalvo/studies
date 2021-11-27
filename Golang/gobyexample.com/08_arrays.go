package main

import "fmt"

func main() {

	//An array is a numbered sequence of elements of a specific lenght
	//Arrays are not commonly used in Go -- slices are used a lot more frequently

	var a [5]int
	fmt.Println("It's empty:", a)

	a[0] = 4
	fmt.Println(a)
	fmt.Println(a[0])
	fmt.Println(len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	var twoD [2][3]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}
