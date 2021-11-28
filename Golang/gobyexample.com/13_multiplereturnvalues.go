package main

import "fmt"

//Go has built in support for multiple return values
//It's an idiomaty thing in go: Return both the result and error values from a function

func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println(a, b)
}
