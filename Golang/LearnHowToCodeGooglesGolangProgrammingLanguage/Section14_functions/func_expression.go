package main

import (
	"fmt"
)

func main() {

	f := func() {
		fmt.Println("My first func expression!")
	}
	f()

	g := func(x int) {
		fmt.Println("Printing the int passed as argument:", x)
	}
	g(99)

	s1 := foo10()
	fmt.Println(s1)

	x := megabar()
	i := x()
	fmt.Println(i)
	//Same things:
	fmt.Println(x())
	fmt.Println(megabar()())

}

func foo10() string {
	return "Hello world"
}

func megabar() func() int {
	return func() int {
		return 451
	}
}
