package main

import (
	"fmt"
	"sort"
)

// Needed to sort in decreasing order
type reverseSort []int

func (a reverseSort) Len() int           { return len(a) }
func (a reverseSort) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a reverseSort) Less(i, j int) bool { return a[i] > a[j] }

func main() {
	//Alrite -- how do I sort this using the sort package?
	a := []int{4, 3, 1, 33, 4, 1, 5}

	sort.Ints(a)
	fmt.Println(a)

	sort.Sort(reverseSort(a))
	fmt.Println(a)
}
