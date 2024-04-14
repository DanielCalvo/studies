package main

import (
	"fmt"
	"golang.org/x/exp/slices"
)

func main() {
	fmt.Println("asd")
	s := []int{1, 1, 1, 6, 14, 5}
	slices.Sort(s)
	fmt.Println(s)                     //Oh my, it is sorted!
	fmt.Println(slices.Contains(s, 1)) //Cool!
	fmt.Println(slices.Compact(s))     //Similar to uniq!

	//Can we use insert to prepend to a slice?
	ss := []int{2, 3, 4, 5}
	ss = slices.Insert(ss, 0, 1) //Inserts 1 at position 0 of the slice
	fmt.Println(ss)              //Will print [1 2 3 4 5]
}
