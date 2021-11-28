package main

import "fmt"

//Go supports pointers! https://en.wikipedia.org/wiki/Pointer_(computer_programming)

func zeroval(n int) {
	n = 0
}

func zeroptr(n *int) {
	*n = 0
}

func main() {

	i := 1
	zeroval(i)
	fmt.Println(i)

	zeroptr(&i)
	fmt.Println(i)

}
