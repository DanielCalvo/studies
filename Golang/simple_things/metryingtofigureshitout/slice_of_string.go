package main

import "fmt"

func main() {
	s := []string{"grape"}
	s = append(s, "banana", "apple")
	fmt.Println(s)

	var ss []string
	ss = append(ss, "grape", "banana", "apple")
}
