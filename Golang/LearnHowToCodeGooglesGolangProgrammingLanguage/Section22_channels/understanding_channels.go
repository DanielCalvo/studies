package main

import "fmt"

func main() {
	c := make(chan int)

	//Channels block!
	go func() {
		c <- 42
	}()

	//This is a buffered channel. It will stop blocking when 1 value is pushed into it.
	cha := make(chan int, 1)
	cha <- 44
	//cha <- 45 //It will break if you try to add a second value to it
	fmt.Println(<-cha)
	fmt.Println(<-c)

	chu := make(chan int, 2) //Instructor recommends to stay away from buffered channels
	chu <- 51
	chu <- 52

	fmt.Println(<-chu, <-chu)

}
