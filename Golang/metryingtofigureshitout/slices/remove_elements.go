package main

import (
	"fmt"
	"golang.org/x/exp/slices"
)

func main() {
	a := []int{0, 1, 2, 3, 4, 5}

	a = removeElements(a, 1, 2)
	a = slices.Compact(a)
	fmt.Println(a)
	fmt.Println(a[1:2])        //To print element at position x: a[x:x+1]
	a = slices.Delete(a, 1, 2) //Deletes element at position 1
	fmt.Println(a)

	//Lets remove the first 2 elements
	a = removeElements(a, 0, 1)
	fmt.Println("A's len:", len(a))
	fmt.Println(a)

	b := []int{1, 2, 3, 4, 5}

	//This will remove element 2, which is "3", but then element 3 will be "5"
	//This outputs leaves you with [1 2 4], instead of [1 2 5]
	b = removeElements(b, 2, 3) //
	fmt.Println(b)
}

// Given a slice and a variadic argument containing positions of elements in the slice to remove them, remove them from the slice and return the slice

// OH GOD you need to sort the element list... and start from the last one? Otherwise removing one element first form the initial array will alter the order of everytHING AAAAAAAAAAA
func removeElements(s []int, elements ...int) []int {

	fmt.Println(elements)
	for _, v := range elements {
		s = slices.Delete(s, v, v+1)
	}
	return s
}
