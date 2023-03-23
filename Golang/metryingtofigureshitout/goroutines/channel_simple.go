package main

import "fmt"

// This is the simplest usage of channels possible I suspect
func main() {

	ch := make(chan int)

	//If I call this without the go keyword, the program deadlocks!
	go sendToChannel(ch)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch) //Receive operations on a close channel will yield the zero value of the channels element type

	fmt.Println("The program finished!")
}

// Let's send three values to a channel!
func sendToChannel(i chan int) chan int {
	i <- 0
	i <- 1
	i <- 2
	close(i) //This channel is now closed. Subsequent attempts to send data to it will cause the program to panic
	return i
}
