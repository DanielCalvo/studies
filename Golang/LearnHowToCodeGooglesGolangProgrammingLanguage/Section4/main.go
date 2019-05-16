package main

import "fmt"

var y = 22

//Declare there is a variable with the identifier z
//And that the variable with the identifier z is of type int
//Assigns the zero value of type int to z
//zero value for its type: false for booleans, 0 for numeric types, "" for strings, and nil for pointers, functions, interfaces, slices, channels, and maps.
var z int

func main() {
	n, err := fmt.Println("Hello world")
	fmt.Println(n, err)

	x := 42
	fmt.Println(x)
	x = 43
	fmt.Println(x)
	//y := 100 + 24
	fmt.Println(y)

	fmt.Println(y)
}

func foo() {
	fmt.Println(y)
}
