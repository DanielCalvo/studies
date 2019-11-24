package main

import (
	"fmt"
	"sync"
	"time"
)

//Notice that the calls to Add are done outside the goroutines they’re helping to track. If
//we didn’t do this, we would have introduced a race condition, because remember
//from “Goroutines” on page 37 that we have no guarantees about when the goroutines
//will be scheduled; we could reach the call to Wait before either of the goroutines
//begin. Had the calls to Add been placed inside the goroutines’ closures, the call to Wait
//could have returned without blocking at all because the calls to Add would not have
//taken place.

func main() {
	var wg sync.WaitGroup

	wg.Add(1) // <1>
	go func() {
		defer wg.Done() // <2>
		fmt.Println("1st goroutine sleeping...")
		time.Sleep(1)
	}()

	wg.Add(1) // <1>
	go func() {
		defer wg.Done() // <2>
		fmt.Println("2nd goroutine sleeping...")
		time.Sleep(2)
	}()

	wg.Wait() // <3>
	fmt.Println("All goroutines complete.")
}
