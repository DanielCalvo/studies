package main

import "fmt"

// This function can only write into the channel
func foo(c chan<- int) {
	c <- 42

}

// This function can only receive from the channel
func bar(c <-chan int) {
	fmt.Println(<-c)

}

func main() {

	c := make(chan int)

	go foo(c)
	bar(c)

	fmt.Println("End of func main")

}
