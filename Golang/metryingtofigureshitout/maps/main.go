package main

import "fmt"

func main() {
	//Lets iterate over a slice and put these elements in a map
	s := []int{1, 2, 3}
	m := make(map[int]bool) //Oh lord, the correct word here is "make" and not "new", hence my confusion...

	for _, v := range s {
		m[v] = true
	}
	fmt.Println(m)
}
