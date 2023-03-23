package main

import "fmt"

var y = 22

// Declare there is a variable with the identifier z
// And that the variable with the identifier z is of type int
// Assigns the zero value of type int to z
// zero value for its type: false for booleans, 0 for numeric types, "" for strings, and nil for pointers, functions, interfaces, slices, channels, and maps.
var z int
var a string = `This is a raw string.
%T "banana"`

type hotdog int

var b hotdog

func main() {
	n, err := fmt.Println("Hello world")
	fmt.Println(n, err)

	//x := 42
	//fmt.Println(x)
	//x = 43
	//fmt.Println(x)
	//y := 100 + 24
	//fmt.Println(y)
	fmt.Printf("Variable y is of type: %T\n", y)
	fmt.Println(a) //Includes line breaks!

	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)
	fmt.Printf("%T\n", b)
	myvar := int(b)
	fmt.Println(myvar)

}

func foo() {
	fmt.Println(y)
}
