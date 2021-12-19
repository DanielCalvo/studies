package main

import "fmt"

func main() {
	mySlice := []string{"a", "b", "c", "d"}

	//Works in slice elements, but not on slice of slices (aka matrixes or multi-dimensional arrays), for those you need to use copy()
	for i := len(mySlice) - 1; i > 0; i-- {
		mySlice[i] = mySlice[i-1]
	}

	mySlice[0] = "E"
	fmt.Println(mySlice)

}
