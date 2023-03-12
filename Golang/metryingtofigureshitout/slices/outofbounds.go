package main

import "fmt"

func main() {
	//s := []string{"a", "b", "c"}
	//fmt.Println(isOutOfBounds(s, 0))
	//fmt.Println(isOutOfBounds(s, 4))

	g := [][]string{
		[]string{"a", "b", "c"},
		[]string{"d", "e", "f"},
		[]string{"g", "h", "i"},
	}
	fmt.Println(g[1][0])
	fmt.Println(len(g), len(g[0]))
	fmt.Println(isOutofBoundsMatrix(g, 3, 2))
}

// Let's assume a square matrix for now
// y is the outer slice
// x is the inner slice
func isOutofBoundsMatrix(s [][]string, y, x int) bool {
	if y < 0 || y > len(s)-1 {
		return true
	}
	//For simplicity, I'm assuming all slices inside the slice have equal length (this matrix is either a square or a rectangle)
	if x < 0 || x > len(s[0])-1 {
		return true
	}
	return false
}

// Oh hey this could be a nice candidate for an interface for matrix/slice
func isOutOfBounds(s []string, pos int) bool {
	if pos < 0 || pos > len(s)-1 { //slices start at 0
		return true
	}
	return false
}
