package main

import (
	"fmt"
)

//Starting with this code, pull the values off the channel using a select statement

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q <-chan int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
	}()

	return c
}

func receive(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println("from the c channel:", v)
		case v := <-q:
			fmt.Println("from the c channel:", v)
		}
	}
}
