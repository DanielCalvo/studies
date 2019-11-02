package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	//send
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(100 * time.Millisecond)
			c <- i
		}
		close(c)

	}()

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("Done!")
}
