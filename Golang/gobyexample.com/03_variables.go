package main

import "fmt"

func main() {
	var a = "inital"
	fmt.Println(a)

	var b, c int = 1, 2 //You can declare multiple variables at once
	fmt.Println(b, c)

	var d = true //Go will infer the type of initialized variables
	fmt.Println(d)

	var e int //Variables declared but not initialized are zero-valued. The zero value of an int is 0
	fmt.Println(e)

	f := "apple" // := is shorthand for declaring and initializing
	fmt.Println(f)
}
