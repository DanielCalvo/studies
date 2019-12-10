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
	go readChan(&wg, c1)
	wg.Wait()

	fmt.Println("End of main")
}

func populateChan(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
		time.Sleep(time.Millisecond * 100)
	}
	close(c)
}

func readChan(wg *sync.WaitGroup, c <-chan int) {
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			defer wg.Done()
			for v := range c {
				fmt.Println("Hello!", v, i)
			}
		}()
	}
}
