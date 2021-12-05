package main

import "fmt"

func main() {
	s := []string{"grape"}
	s = append(s, "banana", "apple")
	fmt.Println(s)

	var ss []string
	ss = append(ss, "grape", "banana", "apple")

	var myslice []string
	myslice = []string{"a", "b", "c"}

	fmt.Println(myslice)
	fmt.Println([]string{"a", "b", "c"})
}
