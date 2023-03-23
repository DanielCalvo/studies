package main

import (
	"fmt"
)

func main() {

	//Example 1
	//Final stage runs on main
	c := gen(2, 3)
	out := sq(c)

	// Consume the output.
	fmt.Println(<-out) // 4
	fmt.Println(<-out) // 9

	//We can also consume the output with a for loop
	for n := range sq(gen(5, 7)) {
		fmt.Println(n) // 16 then 81
	}

}

// First stage
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// Second stage
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
