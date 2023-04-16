package main

import "fmt"

func main() {

	//Iterate from beginning to end of slice, remove one element, iterate all over again

	a := []int{1, 2, 3, 4, 5, 6, 7}

	for len(a) > 0 {
		for i := 0; i < len(a); i++ {
			fmt.Print(a[i])
		}
		fmt.Println()
		a = a[:len(a)-1] //get rid of the last element
	}
}
