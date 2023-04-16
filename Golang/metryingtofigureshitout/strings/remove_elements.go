package main

import (
	"fmt"
	"sort"
)

func main() {
	a := "Hello world!"
	a = removeCharAtIndexes(a, 3, 4)
	a = removeCharAtIndexes(a, 0) //Works with just 1 element too!
	fmt.Println(a)
}

func removeCharAtIndexes(s string, elements ...int) string {
	sort.Ints(elements)
	for i := len(elements) - 1; i >= 0; i-- { //Iterate through elements from end to beginning
		indexToRemove := elements[i] //for better clarity
		s = s[:indexToRemove] + s[indexToRemove+1:]
	}
	return s
}
