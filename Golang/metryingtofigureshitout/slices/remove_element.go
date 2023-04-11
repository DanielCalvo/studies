package main

import "fmt"

func main() {

	i := []int{1, 2, 3, 4, 5}
	s := []string{"1", "2", "3", "4", "5"}

	fmt.Println(RemoveElementFromSlice(i))
	fmt.Println(RemoveElementFromSlice(s))
}

// uh-oh, generics!
// you don't really need a function for this but lets illustrate an example
func RemoveElementFromSlice[T any](slice []T) []T {
	t1 := slice[1:]
	return t1
}
