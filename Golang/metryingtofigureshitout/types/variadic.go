package main

import (
	"fmt"
	"golang.org/x/exp/slices"
)

func main() {
	//Hey whats the type of variadic argument of ints?
	myprint(1, 2, 3)
}

// A variadic param gets converted to a “new” slice inside the func.
// A variadic param is actually syntactic sugar for an input parameter of a slice type.
func myprint(i ...int) {
	//Hey its []int
	fmt.Printf("%T\n", i)

	//Lets try to call some slice operations on it then
	fmt.Println(slices.Contains(i, 1)) //true

	slices.Sort(i)
	fmt.Println(i) //Hey look, decreasing order

}
