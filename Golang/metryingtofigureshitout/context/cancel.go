package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	resultChan := make(chan int)

	go func(ctx context.Context, ch chan<- int) {
		var i int
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutiune cancelled!")
				ch <- i
				return
			default:
				i++
			}
		}
	}(ctx, resultChan)

	time.Sleep(time.Nanosecond)
	fmt.Println("Canceling the goroutine...")
	cancel()

	fmt.Printf("The result is: %d\n", <-resultChan)

}
