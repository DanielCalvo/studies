package main

import (
	"fmt"
)

func main() {

	c := make(chan int)

	go gen(c)
	receive(c)

	fmt.Println("about to exit")
}

func gen(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}

}

func receive(c chan int) {

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
}
