package main

import "fmt"

func main() {

	// Solution 1:
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

	// Solution 2:
	cha := make(chan int, 1)
	cha <- 44
	fmt.Println(<-cha)
}
