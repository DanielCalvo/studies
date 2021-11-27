package main

import "fmt"

func main() {

	s := make([]string, 3)
	fmt.Println(s)
	fmt.Println(len(s))

	s1 := []string{"a", "b", "c"}
	fmt.Println(s1)
	fmt.Println(len(s1))

	s1 = append(s1, "d")
	fmt.Println(s1)

	s2 := make([]string, len(s1))
	copy(s2, s1)
	fmt.Println(s2)

	i := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(i[:3])  //Everything until the element on index 3 of the slice, excluding that element
	fmt.Println(i[3:])  //Everything from element 3 onwards, including the 3 element
	fmt.Println(i[0:3]) //Elements 0, 1 and 2.

}
