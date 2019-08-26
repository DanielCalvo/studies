package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{5, 43, 31, 44, 6, 1, 0, 99}
	sx := []string{"asd", "bbb", "bbc", "aaa", "aa1", "a", "z", "c"}

	sort.Ints(xi)
	fmt.Println(xi)
	sort.Strings(sx)
	fmt.Println(sx)
}
