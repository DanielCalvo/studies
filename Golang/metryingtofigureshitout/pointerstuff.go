package main

import "fmt"

func main() {

	myvar := "banana"
	fmt.Println(myvar)
	fmt.Println(&myvar)

	myvar1 := &myvar
	fmt.Println(*myvar1)
	fmt.Printf("%T", myvar1)
}
