package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := make(chan int)
	var w sync.WaitGroup
	w.Add(5)

	for i := 1; i <= 5; i++ {
		go func(i int, ci <-chan int) {
			for v := range ci {
				time.Sleep(time.Millisecond)
				fmt.Printf("%d got %d\n", i, v)
			}
			w.Done()
		}(i, c)
	}

	for i := 1; i <= 25; i++ {
		c <- i
	}
	close(c)
	w.Wait()
}
