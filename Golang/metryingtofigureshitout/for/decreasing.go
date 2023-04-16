package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{4, 3, 1, 33, 4, 1, 5}
	sort.Ints(a)

	//Classic
	for i := len(a) - 1; i > 0; i-- {
		fmt.Println(a[i])
	}
}
