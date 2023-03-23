package main

import (
	"fmt"
)

// Make a program that counts to 3, sending those values on a channel, but then errors and sends that value to that channel!
func main() {

	//Let's count to three first!
	a := make(chan int)
	e := make(chan error)

	go countToTwo(a, e)

loop:
	for {
		select {
		case num, ok := <-a:
			fmt.Println(num)
			if !ok {
				break loop
			}
		case err := <-e:
			fmt.Println(err)
		}
	}

}

func countToTwo(n chan int, e chan error) {
	n <- 1
	n <- 2
	n <- 2
	n <- 2

	close(n)
	//return n, e

}
