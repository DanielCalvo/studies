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

}
