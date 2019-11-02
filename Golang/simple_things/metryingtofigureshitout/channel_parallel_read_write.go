package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	c1 := make(chan int)

	go readChan(c1)
	populateChan(c1)
	fmt.Println("End of main")

}

func populateChan(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
		time.Sleep(time.Millisecond * 100)
	}
	close(c)
}

func readChan(c chan int) {
	//Not sure if I should have waitgroups in here...
	const goroutines = 8
	for i := 0; i < goroutines; i++ {
		go func() {
			for v := range c {
				fmt.Println(v, runtime.NumGoroutine())
				time.Sleep(time.Millisecond * 300)
			}
		}()

	}
}
