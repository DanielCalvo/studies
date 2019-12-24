package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	c1 := make(chan int)

	var wg sync.WaitGroup
	wg.Add(runtime.NumCPU())

	go populateChan(c1)
	go readChan(c1)

	fmt.Println("End of main")
}

func populateChan(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
		time.Sleep(time.Millisecond * 100)
	}
	close(c)
}

func readChan(c <-chan int) {
	var wg sync.WaitGroup
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			defer wg.Done()
			for v := range c {
				fmt.Println("Hello!", v, i)
			}
		}()
	}
	wg.Wait()
}
